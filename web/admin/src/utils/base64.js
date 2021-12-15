let Base64 = require('js-base64').Base64

// 加密
export function base64Encode(str) {
  return Base64.encode(str) 
}

// 解密
export function base64Decode(str) {
  return Base64.decode(str) 
}