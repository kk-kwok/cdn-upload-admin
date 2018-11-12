const buildType = process.argv.slice(2)[0]
console.log('---------start build--------------' + buildType)

const obj = {
  NODE_ENV: '"production"'
}

switch (buildType) {
  case 'prod':  // 生产
    process.env.srconfig = 'prod'
    obj.srconfig = '"prod"'
    break;
  case 'dev': // 内网开发
    process.env.srconfig = 'dev'
    obj.srconfig = '"dev"'
    break;
  case 'test': // 内网测试
    process.env.srconfig = 'test'
    obj.srconfig = '"test"'
    break;
  default: // 默认开发
    process.env.srconfig = ''
    obj.srconfig = '""'
    break;
}

module.exports = obj;
