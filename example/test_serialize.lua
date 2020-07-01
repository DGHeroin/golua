print('Hello Lua!')

serialize = require('serialize')

local tbl = {
    name = 'my name',
    tick = {
        rate = 60,
        time = 0
    }
}


local address = {
    phone='123',
    address = {
        country='s',
        street='x'
    }
}

-- 序列化成字符串
local str = serialize.encode(tbl, address)

output = io.open("raw.bin", "w")
output:write(str)
output:close()

input = io.open("raw.bin", "r")
str = input:read("*all")
input:close()

print('返回值:', serialize.decode(str))
