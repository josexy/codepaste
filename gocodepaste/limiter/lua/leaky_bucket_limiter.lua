-- ARGV[1]: 最高水位/漏斗最大容量，即最大可以请求数量
-- ARGV[2]: 漏嘴流出水速率（秒），即每秒处理多少个用户请求
-- ARGV[3]: 当前注水时间（秒），即当前用户请求时间

local maxLevel = tonumber(ARGV[1])
local currentVelocity = tonumber(ARGV[2])
local now = tonumber(ARGV[3])

-- 最后访问时间
local lastTime = tonumber(redis.call("hget", KEYS[1], "lastTime"))
-- 当前水位，当前漏斗中水的容量
local curLevel = tonumber(redis.call("hget", KEYS[1], "curLevel"))

-- 1. 初始化
if lastTime == nil then
    lastTime = now
    curLevel = 0
    redis.call("hmset", KEYS[1], "curLevel", curLevel, "lastTime", lastTime)
end

-- 2. 防水，腾出空间

-- 距离上次放水的时间间隔
local interval = now - lastTime
if interval > 0 then
    -- 流出水后新的水位 = 当前水位 - 漏出了多少水
    local newLevel = curLevel - interval * currentVelocity
    if newLevel < 0 then
        newLevel = 0
    end
    curLevel = newLevel
    -- 保存
    redis.call("hmset", KEYS[1], "curLevel", newLevel, "lastTime", now)
end

-- 若到达最高水位，请求失败
if curLevel >= maxLevel then
    return 0
end

-- 3. 注水

-- 反之，若没有到达最高水位，当前水位+1(表示用户成功请求)，请求成功
redis.call("hincrby", KEYS[1], "curLevel", 1)
-- 假设漏斗注满水 / 漏嘴流出水恒速率 = 漏斗中全部水流出的最大时间
-- 即模拟漏斗以恒定速率流出水的过程
redis.call("expire", KEYS[1], maxLevel / currentVelocity)
return 1
