package useragent

import (
	"testing"
)

func TestNormalise(t *testing.T) {

	type test struct {
		UA       string
		Expected string
	}

	tests := []test{

		//removes iOS webview browsers from uastring
		// firefox for iOS
		{"Mozilla/5.0 (iPhone; CPU iPhone OS 12_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/23.0 Mobile/16B92 Safari/605.1.15", "ios_saf/11.0.0"},
		// chrome for iOS
		{"Mozilla/5.0 (iPhone; CPU iPhone OS 12_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/80.0.3987.95 Mobile/15E148 Safari/605.1", "ios_saf/11.0.0"},
		// opera for iOS
		{"Mozilla/5.0 (iPad; CPU OS 11_2_6 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) OPiOS/16.0.8.121059 Mobile/15D100 Safari/9537.53", "ios_saf/11.0.0"},

		// removes Electron browsers from uastring to enable them to report as Chrome
		// Electron for OS X
		{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/537.36 (KHTML, like Gecko) WELLMessenger/1.1.0 Chrome/53.0.2785.143 Electron/1.4.13 Safari/537.36", "chrome/53.0.0"},
		// Electron for Windows
		{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) WELLMessenger/1.1.0 Chrome/53.0.2785.143 Electron/1.4.13 Safari/537.36", "chrome/53.0.0"},

		// removes Facebook in-app browsers from uastring
		// Facebook for iOS
		{"Mozilla/5.0 (iPhone; CPU iPhone OS 9_2 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Mobile/13C75 [FBAN/FBIOS;FBAV/46.0.0.54.156;FBBV/18972819;FBDV/iPhone8,1;FBMD/iPhone;FBSN/iPhone OS;FBSV/9.2;FBSS/2; FBCR/Telenor;FBID/phone;FBLC/nb_NO;FBOP/5]", "ios_saf/9.0.0"},
		// Webview in iOS App with short ua
		{"Mozilla/5.0 (iPhone; CPU iPhone OS 15_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148", "ios_saf/15.2.0"},
		// Facebook for Android, using Chrome browser
		{"Mozilla/5.0 (Linux; Android 4.4.2; SCH-I535 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.0.0 Mobile Safari/537.36 [FBAN/FB4A;FBAV/20.0.0.25.15;]", "chrome/70.0.0"},
	}

	for _, test := range tests {
		client := Normalise(test.UA)

		if client.String() != test.Expected {
			t.Errorf("expected %s, got %s for user agent %s", test.Expected, client.String(), test.UA)
		}
	}
}

func TestAliases(t *testing.T) {

	type test struct {
		UA       string
		Expected string
	}

	tests := []test{
		// Family tests
		// uses browser family name if no alias found
		{"Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.12) Gecko/20101027 Ubuntu/10.04 (lucid) Firefox/50.6.12", "firefox"},
		{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_5; en-us) AppleWebKit/533.18.1 (KHTML, like Gecko) Version/9.0.2 Safari/533.18.5", "safari"},
		{"Mozilla/5.0 (Linux; U; Android 4.3.1; en-us; GT-P7510 Build/HRI83) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13", "android"},
		{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36", "chrome"},

		// uses alias for browser family name if alias exists
		{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:56.0) Gecko/20100101 Firefox/56.0 Waterfox/56.2.12", "firefox"},
		{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.10 Safari/537.36 OPR/27.0.1689.22 (Edition developer)", "chrome"},
		{"Mozilla/5.0 (BB10; Touch) AppleWebKit/537.3+ (KHTML, like Gecko) Version/10.0.9.388 Mobile Safari/537.3+", "bb"},
		{"Mozilla/5.0 (BlackBerry; U; BlackBerry 9930; en) AppleWebKit/534.11+ (KHTML, like Gecko) Version/7.0.0.362 Mobile Safari/534.11+", "bb"},
		{"Mozilla/5.0 (Windows NT 5.1; rv:2.0) Gecko/20110407 Firefox/50.0.3 PaleMoon/50.0.3", "firefox"},
		{"Mozilla/5.0 (Android 5.0; Tablet; rv:41.0) Gecko/41.0 Firefox/41.0", "firefox_mob"},
		{"Mozilla/5.0 (X11; Linux i686 (x86_64); rv:2.0b4) Gecko/20100818 Firefox/45.0b4", "firefox"},
		{"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.3a1) Gecko/20100208 MozillaDeveloperPreview/45.7a1 (.NET CLR 3.5.30729)", "firefox"},
		{"Opera/33.80 (Android 3.2; Linux; Opera Tablet/ADR-1106291546; U; en) Presto/2.8.149 Version/33.10", "chrome"},
		{"Opera/9.80 (S60; SymbOS; Opera Mobi/275; U; es-ES) Presto/2.4.13 Version/10.00", "op_mob"},
		{"SAMSUNG GT-S3330 Opera/9.80 (J2ME/MIDP; Opera Mini/7.1.32840/37.9143; U; en) Presto/2.12.423 Version/12.16", "op_mini"},
		{"Mozilla/5.0 (Linux; Android 4.1.2; GT-S7710 Build/JZO54K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/48.0.1025.166 Mobile", "chrome"},
		{"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; chromeframe/29.0.660.0)", "chrome"},
		{"Mozilla/5.0 (X11; U; Linux i686; en-US) AppleWebKit/534.16 (KHTML, like Gecko) Ubuntu/10.10 Chromium/30.0.648.133 Chrome/30.0.648.133 Safari/534.16", "chrome"},
		{"Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.0; Trident/3.1; IEMobile/11.0; SAMSUNG; SGH-i917)", "ie_mob"},
		{"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; XBLWP7; ZuneWP7)", "ie"},
		{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.3; WOW64; Trident/7.0; .NET4.0E; .NET4.0C; InfoPath.3)", "ie"},
		{"Mozilla/5.0 (Linux; U; Android 2.2.1; en-US; GT-P1000 Build/FROYO) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 UCBrowser/10.0.1.512 U3/0.8.0 Mobile Safari/534.30", "other"},
		{"Mozilla/5.0 (iPhone; CPU iPhone OS 12_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/80.0.3987.95 Mobile/15E148 Safari/605.1", "ios_saf"},
		{"Mozilla/5.0 (iPhone; CPU iPhone OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0 Mobile/15E148 Safari/604.1", "ios_saf"},
		{"Mozilla/5.0 (iPod touch; CPU iPhone OS 9_3_2 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Mobile/13F69", "ios_saf"},
		{"Mozilla/5.0 (iPhone; CPU iPhone OS 9_2 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Mobile/13C75 [FBAN/FBIOS;FBAV/46.0.0.54.156;FBBV/18972819;FBDV/iPhone8,1;FBMD/iPhone;FBSN/iPhone OS;FBSV/9.2;FBSS/2; FBCR/Telenor;FBID/phone;FBLC/nb_NO;FBOP/5]", "ios_saf"},
		{"Mozilla/5.0 (iPhone; CPU iPhone OS 12_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) GSA/66.0.230776083 Mobile/15E148 Safari/605.1", "ios_saf"},
		{"Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16D57 Instagram 80.0.0.12.107", "ios_saf"},
		{"Mozilla/5.0 (Linux; Android 5.0.1; SAMSUNG GT-I9506-ORANGE Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/4.1 Chrome/34.0.1847.76 Mobile Safari/537.36", "samsung_mob"},
		{"Mozilla/5.0 (Macintosh; Intel Mac OS X) AppleWebKit/534.34 (KHTML, like Gecko) PhantomJS/1.6.0 Safari/534.34", "other"},
		{"Mozilla/5.0 (Linux; Android 5.0.1; GT-I9505 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 YaBrowser/14.2.1.1239.00 Mobile Safari/537.36", "chrome"},
		{"Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "chrome"},
		{"Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/66.0.3347.0 Safari/537.36", "chrome"},
	}

	for _, test := range tests {
		client := Normalise(test.UA)

		if client.Family != test.Expected {
			t.Errorf("expected %s, got %s for user agent %s", test.Expected, client.Family, test.UA)
		}
	}
}
