import validator from 'validator'

const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
const rsizes = {
  KB: 1,
  MB: 2,
  GB: 3,
  TB: 4
}
export default {
  bytesToSize(bytes, decimals) {
    if (bytes < 0) return bytes
    if (bytes == 0) return '0 Bytes';
    var k = 1024,
      dm = decimals + 1 || 3,
      sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'],
      i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
  },
  sizeToBytes(size) {
    if (validator.isNumeric(size)) {
      return parseInt(size)
    }
    if (size.length <= 2) {
      return NaN
    }

    let unit = size.slice(-2)
    let pow = rsizes[unit]
    if (pow !== undefined) {
      return parseInt(Math.pow(1024, pow) * parseFloat(size))
    }
    return NaN
  }
}
