package instagram

import "testing"

func Test_getFromMediaPage_image(t *testing.T) {
	caption := "Дружище, есть че?"
	code := "ceiqEstT6r"
	mediaType := "image"
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
	mediaType := "video"
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
