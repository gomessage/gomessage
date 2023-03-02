//判断是否是以数字开头
export function isNumberStart(rule, value, callback) {
  let reg = /^[0-9]+$/
  let check = reg.test(value.toString()); //如果命中这个规则就会返回true
  if (check) {
    callback("不能以数字开头");
  }
  callback();
}

//判断字符串长度
export function isLength(rule, value, callback) {
  if (value.length < 3) {
    callback("长度不能小于3个字符");
  }
  callback();
}


//不能包含特殊字符
export function isStringOrNumber(rule, value, callback) {
  let reg = /^[A-Za-z0-9]+$/
  let check = reg.test(value.toString()); //如果命中这个规则就会返回true
  if (!check) {
    callback("只能是字母或数字，不能包含特殊符号");
  }
  callback();
}
