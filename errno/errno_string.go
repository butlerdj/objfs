// Code generated by "stringer -type=Errno errno.go"; DO NOT EDIT.

package errno

import "strconv"

const _Errno_name = "E2BIGEACCESEADDRINUSEEADDRNOTAVAILEAFNOSUPPORTEAGAINEALREADYEBADFEBADMSGEBUSYECANCELEDECHILDECONNABORTEDECONNREFUSEDECONNRESETEDEADLKEDESTADDRREQEDOMEEXISTEFAULTEFBIGEHOSTUNREACHEIDRMEILSEQEINPROGRESSEINTREINVALEIOEISCONNEISDIRELOOPEMFILEEMLINKEMSGSIZEENAMETOOLONGENETDOWNENETRESETENETUNREACHENFILEENOATTRENOBUFSENODATAENODEVENOENTENOEXECENOLCKENOLINKENOMEMENOMSGENOPROTOOPTENOSPCENOSRENOSTRENOSYSENOTCONNENOTDIRENOTEMPTYENOTRECOVERABLEENOTSOCKENOTSUPENOTTYENXIOEOPNOTSUPPEOVERFLOWEOWNERDEADEPERMEPIPEEPROTOEPROTONOSUPPORTEPROTOTYPEERANGEEROFSESPIPEESRCHETIMEETIMEDOUTETXTBSYEWOULDBLOCKEXDEV"

var _Errno_index = [...]uint16{0, 5, 11, 21, 34, 46, 52, 60, 65, 72, 77, 86, 92, 104, 116, 126, 133, 145, 149, 155, 161, 166, 178, 183, 189, 200, 205, 211, 214, 221, 227, 232, 238, 244, 252, 264, 272, 281, 292, 298, 305, 312, 319, 325, 331, 338, 344, 351, 357, 363, 374, 380, 385, 391, 397, 405, 412, 421, 436, 444, 451, 457, 462, 472, 481, 491, 496, 501, 507, 522, 532, 538, 543, 549, 554, 559, 568, 575, 586, 591}

func (i Errno) String() string {
	i -= 1
	if i < 0 || i >= Errno(len(_Errno_index)-1) {
		return "Errno(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Errno_name[_Errno_index[i]:_Errno_index[i+1]]
}
