import axios from 'axios'

let BaseURL

switch (process.env.srconfig) {
  case 'dev': // 内网开发
    BaseURL = ''
    break
  case 'test': // 内网测试
    BaseURL = ''
    break
  case 'prod': // 生产
    BaseURL = ''
    break
  default: // 默认开发
    BaseURL = ''
    break
}

export {BaseURL, axios}
