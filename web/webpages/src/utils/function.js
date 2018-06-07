import URI from 'urijs'

export default {
  name: 'tool',
  checkData(data) {
    return data === undefined ? '' : data
  },
  rmUIDPrefix(uid) {
    if (uid === undefined) {
      return ''
    }
    let pos = uid.indexOf('-')
    if (pos > 0) {
      return uid.substring(pos + 1, uid.length)
    }
    return uid
  },
  encodeSegment(val) {
    var chars = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'.split(
      ''
    )

    var retr = "";
    var i;
    do{
        i = val % 62
        retr = chars[i] + retr;
        val = (val - i) / 62;
    } while (val > 0)
    return retr;
  },
  decodeSegment(str) {
    var chars = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'.split(
      ''
    )
    var intgers = new Array(128)
    for (var i = 0; i < chars.length; i++) {
      intgers[chars[i]] = i
    }
    var retr = 0
    var len = str.length
    for (var i = 0; i < str.length; i++) {
      retr = retr + intgers[str.charAt(i)] * Math.pow(62, len - i - 1)
    }
    return retr
  },
  leftPad(src, len) {
    src = src.toString()
    if (src == null || src.length >= len) {
      return src
    }
    var padlen = len - src.length
    var retr = src
    for (var i = 0; i < padlen; i++) {
      retr = '0' + retr
    }
    return retr
  },
  encodeBase62(mid){
    var encodeBlockSize = 7
    var decodeBlockSize = 4
    var str = ''
    var midlen = mid.length
    var segments = midlen / encodeBlockSize; // ceil($midlen / $this->encodeBlockSize)
    var start = midlen;
    for (var i = 0; i < segments; i++) {
        start -= encodeBlockSize;
        var seq = mid.substring(start, start + encodeBlockSize);
        seq = this.encodeSegment(parseInt(seq, 10));
        str = this.leftPad(seq, '0', decodeBlockSize) + str;
    }

    if (start > 0) {
        str = this.encodeSegment(parseInt(mid.substring(0, start), 10)) + str;
    }
    return str;
  },
  decodeBase62(str) {
    //解析base62编码的mid
    var encodeBlockSize = 7
    var decodeBlockSize = 4
    var mid = ''
    var strlen = str.length
    var start = strlen
    for (var i = 0; i < parseInt(strlen / decodeBlockSize); i++) {
      start -= decodeBlockSize
      var seg = str.substring(start, start + decodeBlockSize)
      seg = this.decodeSegment(seg)
      mid = this.leftPad(seg, encodeBlockSize) + mid
    }
    if (start > 0) {
      mid = this.decodeSegment(str.substring(0, start)) + mid
    }
    var midLen = mid.length
    var first = mid.charAt(0)
    if (
      (midLen == 16 && (first == '3' || first == '4')) ||
      (midLen == 19 && first == '5')
    ) {
      return mid
    }
    if (mid.charAt(0) == '1' && mid.charAt(7) == '0') {
      mid = mid.substring(0, 7) + mid.substring(8, mid.length())
    }
    return mid
  },
  uidmidFromURL(url) {
    let uri = new URI(url)
    let pathname = uri.pathname()
    if (pathname !== undefined) {
      let ids = pathname.split('/')
      if (ids.length === 3 && ids[1] !== '' && ids[2] !== '') {
        let uid = ids[1]
        let mid = this.decodeBase62(ids[2])
        return { uid: uid, mid: mid }
      }
    }
    return {}
  }
}
