import request from './request';

export function login(userInfo) {
  return request({
    url: '/api/v1/login',
    method: 'post',
    data: userInfo
  });
}
