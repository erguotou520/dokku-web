export function format (date, format) {
  if (!(date instanceof Date)) {
    date = new Date(date)
  }
  return format.replace(/yyyy/, date.getFullYear()).replace(/MM/, date.getMonth() + 1)
    .replace(/dd/, date.getDate()).replace(/hh/, date.getHours()).replace(/mm/, date.getMinutes())
    .replace(/ss/, date.getSeconds())
}
