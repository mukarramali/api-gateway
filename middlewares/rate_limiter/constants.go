package middlewares

const RedisKey = "user-requests:ip"
const Threshold = 30                   // requests per minute
const WindowInMilliseconds = 30 * 1000 // 1 min
