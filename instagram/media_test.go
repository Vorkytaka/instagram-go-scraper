package instagram

import "testing"

func Test_getFromMediaPage_image(t *testing.T) {
	caption := "Дружище, есть че?"
	code := "ceiqEstT6r"
	mediaType := TypeImage
	id := "512999832411258539"
	date := uint64(1375374367)
	ownerID := "248188406"
	ownerUsername := "vorkytaka"
	ownerFullname := "Konstantin"
	likesCount := uint32(33)

	jsonBody := []byte("{\"graphql\":{\"shortcode_media\":{\"__typename\":\"GraphImage\",\"id\":\"512999832411258539\",\"shortcode\":\"ceiqEstT6r\",\"dimensions\":{\"height\":612,\"width\":612},\"gating_info\":null,\"media_preview\":null,\"display_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-15/e15/11325321_112286479105192_62945882_n.jpg\",\"is_video\":false,\"edge_media_to_tagged_user\":{\"edges\":[]},\"edge_media_to_caption\":{\"edges\":[{\"node\":{\"text\":\"Дружище, есть че?\"}}]},\"caption_is_edited\":false,\"edge_media_to_comment\":{\"count\":0,\"page_info\":{\"has_next_page\":false,\"end_cursor\":null},\"edges\":[]},\"comments_disabled\":false,\"taken_at_timestamp\":1375374367,\"edge_media_preview_like\":{\"count\":33,\"edges\":[{\"node\":{\"id\":\"357278100\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/18809075_1893197934237305_7953326533567840256_a.jpg\",\"username\":\"kievaaa_\"}},{\"node\":{\"id\":\"280754522\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/14659324_696211237201737_7228721875210207232_a.jpg\",\"username\":\"vent_blanc\"}},{\"node\":{\"id\":\"264765561\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/18380985_1761898374122867_1200959433912352768_a.jpg\",\"username\":\"scherbakoova\"}},{\"node\":{\"id\":\"272132482\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/11875548_897979613624075_1073258418_a.jpg\",\"username\":\"novanastasiaa\"}},{\"node\":{\"id\":\"43958774\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19228405_353030221779499_5940818490611990528_a.jpg\",\"username\":\"ksksksksksks.ksksks\"}},{\"node\":{\"id\":\"16148732\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19051364_228280957688947_1514116391500775424_a.jpg\",\"username\":\"vladaakim\"}},{\"node\":{\"id\":\"259412979\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19120305_1502589806429988_2144714875857797120_a.jpg\",\"username\":\"asansyy\"}},{\"node\":{\"id\":\"229859976\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/18948064_1936184366653252_4722189294955921408_a.jpg\",\"username\":\"malinyana\"}}]},\"edge_media_to_sponsor_user\":{\"edges\":[]},\"location\":null,\"viewer_has_liked\":false,\"owner\":{\"id\":\"248188406\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19379528_1740774492882462_5097328982882254848_a.jpg\",\"username\":\"vorkytaka\",\"followed_by_viewer\":false,\"full_name\":\"Konstantin\",\"is_private\":false,\"requested_by_viewer\":false,\"is_unpublished\":false,\"blocked_by_viewer\":false,\"has_blocked_viewer\":false},\"is_ad\":false,\"edge_web_media_to_related_media\":{\"edges\":[]}}}}")

	media, err := getFromMediaPage(jsonBody)

	if err != nil {
		t.Error(err)
	}
	if media.Caption != caption {
		t.Errorf("Media caption is incorrect.\nExpect %s, get %s.", caption, media.Caption)
	}
	if media.Code != code || media.MediaList[0].Code != code {
		t.Errorf("Media code is incorrect.\nExpect %s, get %s.", code, media.Code)
	}
	if media.ID != id {
		t.Errorf("Media id is incorrect.\nExpect %s, get %s.", id, media.ID)
	}
	if media.Type != mediaType || media.MediaList[0].Type != mediaType {
		t.Errorf("Media type is incorrect.\nExpect %s, get %s.", mediaType, media.Type)
	}
	if media.Owner.ID != ownerID {
		t.Errorf("Media's owner ID is incorrect.\nExpect %s, get %s.", ownerID, media.Owner.ID)
	}
	if media.Owner.Username != ownerUsername {
		t.Errorf("Media's owner username is incorrect.\nExpect %s, get %s.", ownerUsername, media.Owner.Username)
	}
	if media.Owner.FullName != ownerFullname {
		t.Errorf("Media's owner fullname is incorrect.\nExpect %s, get %s.", ownerFullname, media.Owner.FullName)
	}
	if media.LikesCount != likesCount {
		t.Errorf("Media's likes count is wrong.\nExpect %d, get %d", likesCount, media.LikesCount)
	}
	if media.Date != date {
		t.Errorf("Media's data is wrong.\nExpect %d, get %d", date, media.Date)
	}
}

func Test_getFromMediaPage_video(t *testing.T) {
	caption := "Зов джунглей"
	code := "12376OtT5o"
	mediaType := TypeVideo
	id := "970208779275943528"
	date := uint64(1429877921)
	ownerID := "248188406"
	ownerUsername := "vorkytaka"
	ownerFullname := "Konstantin"
	likesCount := uint32(341)

	jsonBody := []byte("{\"graphql\":{\"shortcode_media\":{\"__typename\":\"GraphVideo\",\"id\":\"970208779275943528\",\"shortcode\":\"12376OtT5o\",\"dimensions\":{\"height\":640,\"width\":640},\"gating_info\":null,\"media_preview\":null,\"display_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-15/e15/11176395_926076444103554_1888764605_n.jpg\",\"video_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t50.2886-16/11175992_1616299585273258_350800542_n.mp4\",\"video_view_count\":0,\"is_video\":true,\"edge_media_to_tagged_user\":{\"edges\":[]},\"edge_media_to_caption\":{\"edges\":[{\"node\":{\"text\":\"Зов джунглей\"}}]},\"caption_is_edited\":false,\"edge_media_to_comment\":{\"count\":4,\"page_info\":{\"has_next_page\":false,\"end_cursor\":null},\"edges\":[{\"node\":{\"id\":\"17851364407060407\",\"text\":\"Король лев😍😍😍\",\"created_at\":1429882239,\"owner\":{\"id\":\"280754522\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/14659324_696211237201737_7228721875210207232_a.jpg\",\"username\":\"vent_blanc\"}}},{\"node\":{\"id\":\"17851364872060407\",\"text\":\"@vent_blanc сынишка Зазу. :)\",\"created_at\":1429900717,\"owner\":{\"id\":\"248188406\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19379528_1740774492882462_5097328982882254848_a.jpg\",\"username\":\"vorkytaka\"}}},{\"node\":{\"id\":\"17851364884060407\",\"text\":\"Тощщнооо))\",\"created_at\":1429901167,\"owner\":{\"id\":\"280754522\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/14659324_696211237201737_7228721875210207232_a.jpg\",\"username\":\"vent_blanc\"}}},{\"node\":{\"id\":\"17874511324035774\",\"text\":\"Хахахаза\",\"created_at\":1488925376,\"owner\":{\"id\":\"259329834\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19367394_565322230522639_6726321571863986176_a.jpg\",\"username\":\"aksinyaren\"}}}]},\"comments_disabled\":false,\"taken_at_timestamp\":1429877921,\"edge_media_preview_like\":{\"count\":341,\"edges\":[{\"node\":{\"id\":\"2201503967\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/17596326_1868782940068419_8276326040075239424_a.jpg\",\"username\":\"numentoy\"}}]},\"edge_media_to_sponsor_user\":{\"edges\":[]},\"location\":null,\"viewer_has_liked\":false,\"owner\":{\"id\":\"248188406\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19379528_1740774492882462_5097328982882254848_a.jpg\",\"username\":\"vorkytaka\",\"followed_by_viewer\":false,\"full_name\":\"Konstantin\",\"is_private\":false,\"requested_by_viewer\":false,\"is_unpublished\":false,\"blocked_by_viewer\":false,\"has_blocked_viewer\":false},\"is_ad\":false,\"edge_web_media_to_related_media\":{\"edges\":[]}}}}")

	media, err := getFromMediaPage(jsonBody)

	if err != nil {
		t.Error(err)
	}
	if media.Caption != caption {
		t.Errorf("Media caption is incorrect.\nExpect %s, get %s.", caption, media.Caption)
	}
	if media.Code != code || media.MediaList[0].Code != code {
		t.Errorf("Media code is incorrect.\nExpect %s, get %s.", code, media.Code)
	}
	if media.ID != id {
		t.Errorf("Media id is incorrect.\nExpect %s, get %s.", id, media.ID)
	}
	if media.Type != mediaType || media.MediaList[0].Type != mediaType {
		t.Errorf("Media type is incorrect.\nExpect %s, get %s.", mediaType, media.Type)
	}
	if media.Owner.ID != ownerID {
		t.Errorf("Media's owner ID is incorrect.\nExpect %s, get %s.", ownerID, media.Owner.ID)
	}
	if media.Owner.Username != ownerUsername {
		t.Errorf("Media's owner username is incorrect.\nExpect %s, get %s.", ownerUsername, media.Owner.Username)
	}
	if media.Owner.FullName != ownerFullname {
		t.Errorf("Media's owner fullname is incorrect.\nExpect %s, get %s.", ownerFullname, media.Owner.FullName)
	}
	if media.LikesCount != likesCount {
		t.Errorf("Media's likes count is wrong.\nExpect %d, get %d", likesCount, media.LikesCount)
	}
	if media.Date != date {
		t.Errorf("Media's data is wrong.\nExpect %d, get %d", date, media.Date)
	}
}

func Test_getFromMediaPage_slider(t *testing.T) {
	caption := "A flamingo-pink parrot tulip. A fistful of sun-yellow ranunculus and magenta anemone. (And not a single bouquet of two dozen long-stemmed red roses in sight.) This is Floom (@floomofficial), an online delivery service that supports locally sourced bouquets. “Everything you’re buying is from an independent florist,” says Lana Elie, the 30-year-old founder of the UK-based business. While working as a personal assistant, she grew frustrated with the process of sending flowers. “Using those companies, it’s the same flowers all year round, but flowers are seasonal. Why do I have to buy roses?” she says. “Flowers are an emotional product — someone’s died, it’s a birthday, a baby is born — and you’re buying from a company that you don’t know and what shows up looks nothing like what you ordered. There’s a real lack of personality.” Luckily, there’s no such problem at Floom. 💐 #MadeToCreate\nPhotos by @floomofficial"
	code := "BUxZSI3B5yq"
	mediaType := TypeCarousel
	id := "1527112946281847978"
	date := uint64(1496266071)
	ownerID := "25025320"
	ownerUsername := "instagram"
	ownerFullname := "Instagram"
	likesCount := uint32(606034)

	mediaListLen := 2
	mediaItemTypes := [2]string{TypeImage, TypeImage}
	mediaItemCodes := [2]string{"BUxZMg1BpMF", "BUxZMgmBhR3"}

	jsonBody := []byte("{\"graphql\":{\"shortcode_media\":{\"__typename\":\"GraphSidecar\",\"id\":\"1527112946281847978\",\"shortcode\":\"BUxZSI3B5yq\",\"dimensions\":{\"height\":1080,\"width\":1080},\"gating_info\":null,\"media_preview\":null,\"display_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-15/e35/18809361_237170043435953_6688961120234897408_n.jpg\",\"is_video\":false,\"edge_media_to_tagged_user\":{\"edges\":[]},\"edge_media_to_caption\":{\"edges\":[{\"node\":{\"text\":\"A flamingo-pink parrot tulip. A fistful of sun-yellow ranunculus and magenta anemone. (And not a single bouquet of two dozen long-stemmed red roses in sight.) This is Floom (@floomofficial), an online delivery service that supports locally sourced bouquets. “Everything you’re buying is from an independent florist,” says Lana Elie, the 30-year-old founder of the UK-based business. While working as a personal assistant, she grew frustrated with the process of sending flowers. “Using those companies, it’s the same flowers all year round, but flowers are seasonal. Why do I have to buy roses?” she says. “Flowers are an emotional product — someone’s died, it’s a birthday, a baby is born — and you’re buying from a company that you don’t know and what shows up looks nothing like what you ordered. There’s a real lack of personality.” Luckily, there’s no such problem at Floom. 💐 #MadeToCreate\\nPhotos by @floomofficial\"}}]},\"caption_is_edited\":false,\"edge_media_to_comment\":{\"count\":2746,\"page_info\":{\"has_next_page\":true,\"end_cursor\":\"AQAO1E08FD0fKn_9pVXOzsrXeThniS3rT3jNEU4ppvUM8ksUxRI3aEuTIAD1tVBH3Ak0jrsxQ39CPcySMpa0mQe-H4j2izXnoy9VWQ8_tRXCaw\"},\"edges\":[{\"node\":{\"id\":\"17859486364162686\",\"text\":\"omg\",\"created_at\":1496924001,\"owner\":{\"id\":\"1648829755\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19534463_1287878458001364_4861927215572451328_a.jpg\",\"username\":\"coolboy_santosh\"}}},{\"node\":{\"id\":\"17871744292102520\",\"text\":\"Beautiful\",\"created_at\":1497018589,\"owner\":{\"id\":\"2126717422\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19122178_1695163520778717_945265230324170752_a.jpg\",\"username\":\"lionel_yogi\"}}},{\"node\":{\"id\":\"17868865576091472\",\"text\":\"Vauuuu\",\"created_at\":1497031140,\"owner\":{\"id\":\"4665779593\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/18888735_617200658480500_2382322916327948288_a.jpg\",\"username\":\"alievaaa.___\"}}},{\"node\":{\"id\":\"17874701056073341\",\"text\":\"beautiful\",\"created_at\":1497088043,\"owner\":{\"id\":\"4916449111\",\"profile_pic_url\":\"https://scontent-scl1-1.cdninstagram.com/t51.2885-19/11906329_960233084022564_1448528159_a.jpg\",\"username\":\"vananasufi\"}}},{\"node\":{\"id\":\"17939511772019133\",\"text\":\"Beautiful.\",\"created_at\":1497104051,\"owner\":{\"id\":\"4821382811\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/17267416_757995227701223_6768796732542156800_a.jpg\",\"username\":\"palsatya276\"}}},{\"node\":{\"id\":\"17883878596029303\",\"text\":\"🎈🎈🎈\",\"created_at\":1497167153,\"owner\":{\"id\":\"4387966109\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19623180_439092366475166_2083033918512037888_a.jpg\",\"username\":\"_s_afa\"}}},{\"node\":{\"id\":\"17883892915024944\",\"text\":\"Pretty..\",\"created_at\":1497193057,\"owner\":{\"id\":\"365180161\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19379565_1362394680512584_7966230930312396800_a.jpg\",\"username\":\"pieyqasso_\"}}},{\"node\":{\"id\":\"17859937681150664\",\"text\":\"Pretty 👍🏻\",\"created_at\":1497205371,\"owner\":{\"id\":\"2311282377\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19050305_2048126058733130_2326962437149949952_n.jpg\",\"username\":\"angela11783\"}}},{\"node\":{\"id\":\"17859513400157880\",\"text\":\"Gorgeous post!\",\"created_at\":1497290183,\"owner\":{\"id\":\"4304183836\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/16465606_1875120946066687_146962064198336512_a.jpg\",\"username\":\"amadeuseditorial\"}}},{\"node\":{\"id\":\"17869501435083866\",\"text\":\"زیبا\",\"created_at\":1497307746,\"owner\":{\"id\":\"2209923988\",\"profile_pic_url\":\"https://scontent-scl1-1.cdninstagram.com/t51.2885-19/11906329_960233084022564_1448528159_a.jpg\",\"username\":\"emil.dost\"}}},{\"node\":{\"id\":\"17860240501145077\",\"text\":\"💖💖\",\"created_at\":1497369761,\"owner\":{\"id\":\"3556057829\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/14716391_1192815317428973_1878546548409761792_a.jpg\",\"username\":\"thaisbriitos\"}}},{\"node\":{\"id\":\"17844028333197282\",\"text\":\"It's so so beautiful and lovely 😍😍❤\",\"created_at\":1497373978,\"owner\":{\"id\":\"3746888431\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19428663_1468426889885187_376592086952050688_a.jpg\",\"username\":\"elenacarcea\"}}},{\"node\":{\"id\":\"17878107634066995\",\"text\":\"@azerreza1246 😉\",\"created_at\":1497383520,\"owner\":{\"id\":\"3619070416\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/18809110_131531610748628_3114607109354815488_a.jpg\",\"username\":\"ma4sharu_\"}}},{\"node\":{\"id\":\"17884077922011137\",\"text\":\"@krn_shekhawat ❤\",\"created_at\":1497383528,\"owner\":{\"id\":\"3619070416\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/18809110_131531610748628_3114607109354815488_a.jpg\",\"username\":\"ma4sharu_\"}}},{\"node\":{\"id\":\"17871221851123716\",\"text\":\"@sergiocarcamosoto )\",\"created_at\":1497383556,\"owner\":{\"id\":\"3619070416\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/18809110_131531610748628_3114607109354815488_a.jpg\",\"username\":\"ma4sharu_\"}}},{\"node\":{\"id\":\"17871009106122753\",\"text\":\"@danielakoha ❤\",\"created_at\":1497383572,\"owner\":{\"id\":\"3619070416\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/18809110_131531610748628_3114607109354815488_a.jpg\",\"username\":\"ma4sharu_\"}}},{\"node\":{\"id\":\"17885210545057777\",\"text\":\"I agree with you.\",\"created_at\":1497406386,\"owner\":{\"id\":\"4461394034\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/17493950_1618933134801860_3033484441377832960_a.jpg\",\"username\":\"1wrecklessblonde\"}}},{\"node\":{\"id\":\"17858456071174687\",\"text\":\"J'aime beaucoup\",\"created_at\":1497457190,\"owner\":{\"id\":\"4023282767\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/14596602_130857864045231_7083362104396218368_a.jpg\",\"username\":\"nassibnabli\"}}},{\"node\":{\"id\":\"17883834886057724\",\"text\":\"Wow\",\"created_at\":1497460651,\"owner\":{\"id\":\"3879200094\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19428897_149222422291756_5017979217275518976_a.jpg\",\"username\":\"attitude_____wla_munda________\"}}},{\"node\":{\"id\":\"17844156853196884\",\"text\":\"Beautifulll\",\"created_at\":1497470075,\"owner\":{\"id\":\"2233664386\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/18947909_173413103192292_8907986519520706560_a.jpg\",\"username\":\"hawaninmutfagi\"}}},{\"node\":{\"id\":\"17860417675150505\",\"text\":\"Woooh j adore\",\"created_at\":1497507217,\"owner\":{\"id\":\"3953915376\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/14334251_154939141624381_1399011957_a.jpg\",\"username\":\"simone.francone\"}}},{\"node\":{\"id\":\"17883765469057938\",\"text\":\"May be too bright\",\"created_at\":1497535935,\"owner\":{\"id\":\"5567566772\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/18888965_477263569282571_8011009155533897728_a.jpg\",\"username\":\"3pteb\"}}},{\"node\":{\"id\":\"17869936552082471\",\"text\":\"Beautiful\",\"created_at\":1497719895,\"owner\":{\"id\":\"2308579230\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/13704182_283206725368307_1929270224_a.jpg\",\"username\":\"susang354\"}}},{\"node\":{\"id\":\"17884104046031176\",\"text\":\"💟\",\"created_at\":1497737084,\"owner\":{\"id\":\"1828269560\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19227807_1528079490565110_6343451703014064128_a.jpg\",\"username\":\"mrs.bugzy\"}}},{\"node\":{\"id\":\"17884552093046978\",\"text\":\"NoordinKhan  9050768732\",\"created_at\":1497741416,\"owner\":{\"id\":\"5613269899\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19120488_255880688223245_8264418015613812736_a.jpg\",\"username\":\"noordin.khan.16100\"}}},{\"node\":{\"id\":\"17858978443186407\",\"text\":\"Noordin 786\",\"created_at\":1497741477,\"owner\":{\"id\":\"5613269899\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19120488_255880688223245_8264418015613812736_a.jpg\",\"username\":\"noordin.khan.16100\"}}},{\"node\":{\"id\":\"17871713965122315\",\"text\":\"Очень красиво\",\"created_at\":1497758364,\"owner\":{\"id\":\"4781836179\",\"profile_pic_url\":\"https://scontent-scl1-1.cdninstagram.com/t51.2885-19/11906329_960233084022564_1448528159_a.jpg\",\"username\":\"356903_\"}}},{\"node\":{\"id\":\"17872390432118535\",\"text\":\"Ok\",\"created_at\":1497768764,\"owner\":{\"id\":\"4735975059\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/16906605_220053978468159_8746759455888113664_a.jpg\",\"username\":\"arnoldlewis26\"}}},{\"node\":{\"id\":\"17870045152093465\",\"text\":\"F\",\"created_at\":1497787315,\"owner\":{\"id\":\"3851721815\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/18723512_1514298181974890_889975910164332544_a.jpg\",\"username\":\"ismael_aliraqe\"}}},{\"node\":{\"id\":\"17858832070193334\",\"text\":\"Nice\",\"created_at\":1497906271,\"owner\":{\"id\":\"2986419488\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19367065_1965715420327387_6738447067699478528_a.jpg\",\"username\":\"seckin_sevinc\"}}},{\"node\":{\"id\":\"17885683252049371\",\"text\":\"Beautiful 😍😍😍\",\"created_at\":1498054069,\"owner\":{\"id\":\"4547595465\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19121613_1881063362153485_1651875027978551296_a.jpg\",\"username\":\"hesam777242\"}}},{\"node\":{\"id\":\"17870804005082687\",\"text\":\"Nice\",\"created_at\":1498193151,\"owner\":{\"id\":\"5412346167\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/18161782_134813537062032_6744477206078029824_a.jpg\",\"username\":\"ajitkumarpradhan942.akp59\"}}},{\"node\":{\"id\":\"17873122642116287\",\"text\":\"Nicepice\",\"created_at\":1498210133,\"owner\":{\"id\":\"2541979466\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/12383270_210205625989435_1977010348_a.jpg\",\"username\":\"vipin.kumar5445\"}}},{\"node\":{\"id\":\"17874422794107709\",\"text\":\"Nc\",\"created_at\":1498751563,\"owner\":{\"id\":\"5652159909\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19436777_213119295876763_394536486084542464_a.jpg\",\"username\":\"dennxxyox\"}}}]},\"comments_disabled\":false,\"taken_at_timestamp\":1496266071,\"edge_media_preview_like\":{\"count\":606034,\"edges\":[{\"node\":{\"id\":\"5572234202\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19437076_1399720146776347_4728347397590089728_a.jpg\",\"username\":\"sintyaprmt_\"}},{\"node\":{\"id\":\"5456893376\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19534860_442176156163808_7988128413049159680_a.jpg\",\"username\":\"fatiha.din\"}},{\"node\":{\"id\":\"5605282646\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19120267_268080040266766_96546625681358848_a.jpg\",\"username\":\"natalisiyahbayaz4\"}},{\"node\":{\"id\":\"4202903717\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19436388_143461062872534_2857203478718054400_a.jpg\",\"username\":\"pangestutobi1\"}},{\"node\":{\"id\":\"1572532039\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/10707051_697407970356729_2108764909_a.jpg\",\"username\":\"deek206\"}},{\"node\":{\"id\":\"4229785793\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/15875875_198386170630208_4114841345472856064_n.jpg\",\"username\":\"ine808\"}},{\"node\":{\"id\":\"4940966673\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/17268045_353091315085908_8690900584410120192_a.jpg\",\"username\":\"lauradominguez8321\"}},{\"node\":{\"id\":\"3163630920\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19533723_1883318508588589_6895443652970020864_a.jpg\",\"username\":\"alfira_sha\"}},{\"node\":{\"id\":\"2165527002\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/18879285_749063495267087_3214072564745764864_a.jpg\",\"username\":\"putrimysyrl_\"}},{\"node\":{\"id\":\"4072126623\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/14624773_192402567830574_8511361107372277760_a.jpg\",\"username\":\"zambranofica\"}}]},\"edge_media_to_sponsor_user\":{\"edges\":[]},\"location\":null,\"viewer_has_liked\":false,\"owner\":{\"id\":\"25025320\",\"profile_pic_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/14719833_310540259320655_1605122788543168512_a.jpg\",\"username\":\"instagram\",\"followed_by_viewer\":false,\"full_name\":\"Instagram\",\"is_private\":false,\"requested_by_viewer\":false,\"is_unpublished\":false,\"blocked_by_viewer\":false,\"has_blocked_viewer\":false},\"is_ad\":false,\"edge_web_media_to_related_media\":{\"edges\":[]},\"edge_sidecar_to_children\":{\"edges\":[{\"node\":{\"__typename\":\"GraphImage\",\"id\":\"1527112559701168901\",\"shortcode\":\"BUxZMg1BpMF\",\"dimensions\":{\"height\":1080,\"width\":1080},\"gating_info\":null,\"media_preview\":\"ACoq6OmswUZPAFKTis2W44OTwRjI6dffpxwaluxRobqdms+Jh8oY5UcjryfU+w7VYaXBAx2zk9KE0xalmimqadTGRStt61hzu6uW6DHI+vA/Pr9a25lyM4zjqPUVUVVY7mAPrnsO34D1pMNiC3Df6wj6gHkHAzx0x9Pxq1IDKMKcYOTnuPb+tMiwPmi5Bzkjt+H/ANbpVpY1z06iotcLj1APIJ44x2qSmqoXoMZ5p1aCQVVktgeV/KrVFAyjHZojbgCCPc4/KroFLRQKwUUUUDP/2Q==\",\"display_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-15/e35/18809361_237170043435953_6688961120234897408_n.jpg\",\"is_video\":false,\"edge_media_to_tagged_user\":{\"edges\":[]}}},{\"node\":{\"__typename\":\"GraphImage\",\"id\":\"1527112559449478263\",\"shortcode\":\"BUxZMgmBhR3\",\"dimensions\":{\"height\":1080,\"width\":1080},\"gating_info\":null,\"media_preview\":\"ACoq6Oo3lVOp6daZPJ5a8dTwP8fwrJ8wknH0bnnufQ9aluxRsLJmpM1mJMWACHHPJI6Drj3yOlRXHm7gTlk7Y7fXHf8AzmlzA1Y2aKih3bfmx+FS1QFC95KjtzWPG4eTC8EdeByOn4euevSt+4j3rnrisl4XXDEAeo7+/HvxUSBFgHew7gDOOhB5H49+mfrU8rBYiX5/zxWfGGdtyDgYzjqD9CenH1/OtH7xAb/9RqLFN6E0LqUXHyg8Af571YqKOPZxnP8An/PNS1sQFQvFuHFTUUDKgttpypwT1qyBTqKLWAKKKKAP/9k=\",\"display_url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-15/e35/18809449_1911468069133594_8666172859959214080_n.jpg\",\"is_video\":false,\"edge_media_to_tagged_user\":{\"edges\":[]}}}]}}}}")

	media, err := getFromMediaPage(jsonBody)

	if err != nil {
		t.Error(err)
	}
	if media.Caption != caption {
		t.Errorf("Media caption is incorrect.\nExpect %s, get %s.", caption, media.Caption)
	}
	if media.Code != code {
		t.Errorf("Media code is incorrect.\nExpect %s, get %s.", code, media.Code)
	}
	if media.ID != id {
		t.Errorf("Media id is incorrect.\nExpect %s, get %s.", id, media.ID)
	}
	if media.Type != mediaType {
		t.Errorf("Media type is incorrect.\nExpect %s, get %s.", mediaType, media.Type)
	}
	if media.Owner.ID != ownerID {
		t.Errorf("Media's owner ID is incorrect.\nExpect %s, get %s.", ownerID, media.Owner.ID)
	}
	if media.Owner.Username != ownerUsername {
		t.Errorf("Media's owner username is incorrect.\nExpect %s, get %s.", ownerUsername, media.Owner.Username)
	}
	if media.Owner.FullName != ownerFullname {
		t.Errorf("Media's owner fullname is incorrect.\nExpect %s, get %s.", ownerFullname, media.Owner.FullName)
	}
	if media.LikesCount != likesCount {
		t.Errorf("Media's likes count is wrong.\nExpect %d, get %d", likesCount, media.LikesCount)
	}
	if media.Date != date {
		t.Errorf("Media's data is wrong.\nExpect %d, get %d", date, media.Date)
	}
	if len(media.MediaList) != mediaListLen {
		t.Errorf("Incorrect media list length.\nExpect %d, get %d", mediaListLen, len(media.MediaList))
	}
	for i, m := range media.MediaList {
		if m.Code != mediaItemCodes[i] {
			t.Errorf("Media code is incorrect.\nExpect %s, get %s.", mediaItemCodes[i], m.Code)
		}
		if m.Type != mediaItemTypes[i] {
			t.Errorf("Media type is incorrect.\nExpect %s, get %s.", mediaItemTypes[i], m.Type)
		}
	}
}

func Test_getFromAccountMediaList(t *testing.T) {
	caption := "Всем добра и расслабона."
	code := "YgBPkNtTz_"
	mediaType := TypeImage
	id := "441358231205657855_248188406"
	date := uint64(1366834022)
	ownerID := "248188406"
	ownerUsername := "vorkytaka"
	ownerFullname := "Konstantin"
	likesCount := uint32(26)

	jsonBody := []byte("{\"id\":\"441358231205657855_248188406\",\"code\":\"YgBPkNtTz_\",\"user\":{\"id\":\"248188406\",\"full_name\":\"Konstantin\",\"profile_picture\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19379528_1740774492882462_5097328982882254848_a.jpg\",\"username\":\"vorkytaka\"},\"images\":{\"thumbnail\":{\"width\":150,\"height\":150,\"url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-15/s150x150/e15/11330527_1587491721510550_960812045_n.jpg\"},\"low_resolution\":{\"width\":320,\"height\":320,\"url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-15/s320x320/e15/11330527_1587491721510550_960812045_n.jpg\"},\"standard_resolution\":{\"width\":612,\"height\":612,\"url\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-15/e15/11330527_1587491721510550_960812045_n.jpg\"}},\"created_time\":\"1366834022\",\"caption\":{\"id\":\"17850024664060407\",\"text\":\"Всем добра и расслабона.\",\"created_time\":\"1366834022\",\"from\":{\"id\":\"248188406\",\"full_name\":\"Konstantin\",\"profile_picture\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19379528_1740774492882462_5097328982882254848_a.jpg\",\"username\":\"vorkytaka\"}},\"user_has_liked\":false,\"likes\":{\"data\":[{\"id\":\"278425203\",\"full_name\":\"Fahri Muradov\",\"profile_picture\":\"https://scontent-cdt1-1.cdninstagram.com/t51.2885-19/11906329_960233084022564_1448528159_a.jpg\",\"username\":\"fahrimuradov\"},{\"id\":\"23777215\",\"full_name\":\"Антон\",\"profile_picture\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/13628374_292016351178975_1594723130_a.jpg\",\"username\":\"anton.checherin\"},{\"id\":\"259412979\",\"full_name\":\"⠀⠀⠀⠀⠀⠀⠀⠀⠀➰ Anastasia Feduró ➰\",\"profile_picture\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19120305_1502589806429988_2144714875857797120_a.jpg\",\"username\":\"asansyy\"},{\"id\":\"12910777\",\"full_name\":\"Yury Isaev\",\"profile_picture\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/11055431_859946827394988_1161718002_a.jpg\",\"username\":\"basay17\"}],\"count\":26},\"comments\":{\"data\":[{\"id\":\"17846804761060407\",\"text\":\"Это уже перебор D:\",\"created_time\":\"1366874266\",\"from\":{\"id\":\"215969458\",\"full_name\":\"Oleksandr Zahorulia\",\"profile_picture\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/13188116_1715562422017072_2142692244_a.jpg\",\"username\":\"hast4656\"}},{\"id\":\"17846806420060407\",\"text\":\"@hast4656 ты еще первый вариант не видел... :(\",\"created_time\":\"1367136817\",\"from\":{\"id\":\"248188406\",\"full_name\":\"Konstantin\",\"profile_picture\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/19379528_1740774492882462_5097328982882254848_a.jpg\",\"username\":\"vorkytaka\"}},{\"id\":\"17848842157060407\",\"text\":\"Ахаха))\",\"created_time\":\"1390122094\",\"from\":{\"id\":\"210783735\",\"full_name\":\"Константин Сахно\",\"profile_picture\":\"https://instagram.fhen1-1.fna.fbcdn.net/t51.2885-19/s150x150/13739563_270619919990094_445324022_a.jpg\",\"username\":\"sahno.life\"}}],\"count\":8},\"can_view_comments\":true,\"can_delete_comments\":true,\"type\":\"image\",\"link\":\"https://www.instagram.com/p/YgBPkNtTz_/\",\"location\":null,\"alt_media_url\":null}")

	media, err := getFromAccountMediaList(jsonBody)

	if err != nil {
		t.Error(err)
	}
	if media.Caption != caption {
		t.Errorf("Media caption is incorrect.\nExpect %s, get %s.", caption, media.Caption)
	}
	if media.Code != code {
		t.Errorf("Media code is incorrect.\nExpect %s, get %s.", code, media.Code)
	}
	if media.ID != id {
		t.Errorf("Media id is incorrect.\nExpect %s, get %s.", id, media.ID)
	}
	if media.Type != mediaType {
		t.Errorf("Media type is incorrect.\nExpect %s, get %s.", mediaType, media.Type)
	}
	if media.Owner.ID != ownerID {
		t.Errorf("Media's owner ID is incorrect.\nExpect %s, get %s.", ownerID, media.Owner.ID)
	}
	if media.Owner.Username != ownerUsername {
		t.Errorf("Media's owner username is incorrect.\nExpect %s, get %s.", ownerUsername, media.Owner.Username)
	}
	if media.Owner.FullName != ownerFullname {
		t.Errorf("Media's owner fullname is incorrect.\nExpect %s, get %s.", ownerFullname, media.Owner.FullName)
	}
	if media.LikesCount != likesCount {
		t.Errorf("Media's likes count is wrong.\nExpect %d, get %d", likesCount, media.LikesCount)
	}
	if media.Date != date {
		t.Errorf("Media's data is wrong.\nExpect %d, get %d", date, media.Date)
	}
}