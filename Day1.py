import math

def getFuelReq(mass):
	return math.floor(mass / 3) - 2

def getFuelReqNew(mass):
	fuelReq = 0
	while ((math.floor(mass / 3) - 2) > 0):
		fuelReq += math.floor(mass / 3) - 2
		mass = math.floor(mass / 3) - 2
	return fuelReq

totalFuel = \
getFuelReq(53247) + \
getFuelReq(140268) + \
getFuelReq(139961) + \
getFuelReq(87816) + \
getFuelReq(102986) + \
getFuelReq(122415) + \
getFuelReq(140484) + \
getFuelReq(56099) + \
getFuelReq(52832) + \
getFuelReq(56999) + \
getFuelReq(122823) + \
getFuelReq(130608) + \
getFuelReq(83149) + \
getFuelReq(144224) + \
getFuelReq(104559) + \
getFuelReq(108523) + \
getFuelReq(126571) + \
getFuelReq(137284) + \
getFuelReq(83197) + \
getFuelReq(109996) + \
getFuelReq(56795) + \
getFuelReq(73112) + \
getFuelReq(50043) + \
getFuelReq(130097) + \
getFuelReq(113563) + \
getFuelReq(91200) + \
getFuelReq(130589) + \
getFuelReq(83725) + \
getFuelReq(108625) + \
getFuelReq(131977) + \
getFuelReq(95213) + \
getFuelReq(149800) + \
getFuelReq(70756) + \
getFuelReq(86240) + \
getFuelReq(134854) + \
getFuelReq(148919) + \
getFuelReq(114460) + \
getFuelReq(95062) + \
getFuelReq(122989) + \
getFuelReq(57274) + \
getFuelReq(112074) + \
getFuelReq(139530) + \
getFuelReq(131217) + \
getFuelReq(55652) + \
getFuelReq(125522) + \
getFuelReq(77304) + \
getFuelReq(144873) + \
getFuelReq(86811) + \
getFuelReq(107768) + \
getFuelReq(70704) + \
getFuelReq(146300) + \
getFuelReq(87256) + \
getFuelReq(118752) + \
getFuelReq(52585) + \
getFuelReq(126269) + \
getFuelReq(124559) + \
getFuelReq(52592) + \
getFuelReq(112097) + \
getFuelReq(123090) + \
getFuelReq(101778) + \
getFuelReq(136424) + \
getFuelReq(74547) + \
getFuelReq(119337) + \
getFuelReq(143570) + \
getFuelReq(113426) + \
getFuelReq(146458) + \
getFuelReq(88135) + \
getFuelReq(100236) + \
getFuelReq(148224) + \
getFuelReq(98718) + \
getFuelReq(135181) + \
getFuelReq(56608) + \
getFuelReq(109671) + \
getFuelReq(144027) + \
getFuelReq(135192) + \
getFuelReq(111620) + \
getFuelReq(69411) + \
getFuelReq(107957) + \
getFuelReq(88448) + \
getFuelReq(64972) + \
getFuelReq(63010) + \
getFuelReq(100625) + \
getFuelReq(96144) + \
getFuelReq(61998) + \
getFuelReq(59813) + \
getFuelReq(124503) + \
getFuelReq(64306) + \
getFuelReq(73119) + \
getFuelReq(77094) + \
getFuelReq(136295) + \
getFuelReq(132224) + \
getFuelReq(125713) + \
getFuelReq(110137) + \
getFuelReq(51478) + \
getFuelReq(90272) + \
getFuelReq(133506) + \
getFuelReq(72218) + \
getFuelReq(100082) + \
getFuelReq(106377) + \
getFuelReq(140290)

print(totalFuel)
