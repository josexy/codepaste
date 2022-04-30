-- ARGV[1]: 窗口时间大小
-- ARGV[2]: 窗口请求上限
-- ARGV[3]: 当前小窗口值
-- ARGV[4]: 起始小窗口值

local window = tonumber(ARGV[1])
local limit = tonumber(ARGV[2])
local currentSmallWindow = tonumber(ARGV[3])
local startSmallWindow = tonumber(ARGV[4])

-- 计算当前窗口的请求总数
local counters = redis.call("hgetall", KEYS[1])
local count = 0

-- 统计所有小窗口计数器
for i = 1, #(counters) / 2 do
    local smallWindow = tonumber(counters[i * 2 - 1])
    local counter = tonumber(counters[i * 2])
    if smallWindow < startSmallWindow then
        redis.call("hdel", KEYS[1], smallWindow)
    else
        count = count + counter
    end
end

-- 若到达窗口请求上限，请求失败
if count >= limit then
    return 0
end

-- 若没到窗口请求上限，当前小窗口计数器+1，请求成功
redis.call("hincrby", KEYS[1], currentSmallWindow, 1)
redis.call("pexpire", KEYS[1], window)
return 1