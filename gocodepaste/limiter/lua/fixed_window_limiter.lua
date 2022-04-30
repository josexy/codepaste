-- ARGV[1]: 窗口时间大小
-- ARGV[2]: 窗口请求上限

local window = tonumber(ARGV[1])
local limit = tonumber(ARGV[2])

-- 获取原始值
local counter = tonumber(redis.call("get", KEYS[1]))
if counter == nil then
    counter = 0
end
-- 若到达窗口请求上限，请求失败
if counter >= limit then
    return 0
end
-- 窗口值+1
redis.call("incr", KEYS[1])
-- 设置窗口过期时间，单位毫秒!!! 不同于 expire
if counter == 0 then
    redis.call("pexpire", KEYS[1], window)
end
return 1