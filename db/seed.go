package db

import (
	dbModel "backend-api/model/db"
	"backend-api/utils/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RawQuiz struct {
	Title     string
	Questions []RawQuestion
}

type RawQuestion struct {
	Text    string
	Options []string
	Answer  int
}

func SeedQuizzesFullSet(c *gin.Context) error {
	var existing *dbModel.Quiz
	err := DB.First(&existing).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err != nil {
			logger.InternalServerError(c, err)
			return err
		}
		if existing != nil {
			logger.BadRequest(c, "already seeded")
			return nil
		}
	}

	quizzes := []RawQuiz{
		{
			Title:     "Tiếng Anh cơ bản",
			Questions: genEnglishQuestions(),
		},
		{
			Title:     "Công nghệ thông tin",
			Questions: genITQuestions(),
		},
		{
			Title:     "IQ logic",
			Questions: genIQQuestions(),
		},
		{
			Title:     "Văn học Việt Nam",
			Questions: genLiteratureQuestions(),
		},
		{
			Title:     "Âm nhạc",
			Questions: genMusicQuestions(),
		},
		{
			Title:     "Nghệ thuật",
			Questions: genArtQuestions(),
		},
		{
			Title:     "Thể thao",
			Questions: genSportQuestions(),
		},
		{
			Title:     "Môi trường",
			Questions: genEnvironmentQuestions(),
		},
		{
			Title:     "Văn hóa thế giới",
			Questions: genWorldCultureQuestions(),
		},
		{
			Title:     "Đạo đức học",
			Questions: genEthicsQuestions(),
		},
		{
			Title:     "Kỹ năng sống",
			Questions: genLifeSkillQuestions(),
		},
		{
			Title:     "Khoa học xã hội",
			Questions: genSocialScienceQuestions(),
		},
		{
			Title:     "Tiếng Việt",
			Questions: genVietnameseQuestions(),
		},
		{
			Title:     "Lập trình cơ bản",
			Questions: genCodingQuestions(),
		},
		{
			Title:     "Toán học cơ bản",
			Questions: genMathQuestions(),
		},
		{
			Title:     "Khoa học tự nhiên",
			Questions: genScienceQuestions(),
		},
	}

	for _, raw := range quizzes {
		quiz := dbModel.Quiz{Title: raw.Title}
		if err := DB.Create(&quiz).Error; err != nil {
			logger.InternalServerError(c, err)
			return err
		}
		for _, q := range raw.Questions {
			question := dbModel.Question{QuizID: quiz.ID, QuestionText: q.Text, Score: 1}
			if err := DB.Create(&question).Error; err != nil {
				logger.InternalServerError(c, err)
				return err
			}
			for i, opt := range q.Options {
				ans := dbModel.AnswerOption{
					QuestionID: question.ID,
					Text:       opt,
					IsCorrect:  i == q.Answer,
				}
				if err := DB.Create(&ans).Error; err != nil {
					logger.InternalServerError(c, err)
					return err
				}
			}
		}
	}
	return nil
}

func genEnglishQuestions() []RawQuestion {
	return []RawQuestion{
		{"'Hello' nghĩa là gì?", []string{"Xin chào", "Tạm biệt", "Cảm ơn", "Làm ơn"}, 0},
		{"'Apple' là gì?", []string{"Quả táo", "Quả chuối", "Quả nho", "Quả lê"}, 0},
		{"Số 'five' là số mấy?", []string{"5", "4", "6", "3"}, 0},
		{"'Thank you' có nghĩa là?", []string{"Cảm ơn", "Xin lỗi", "Không", "Có"}, 0},
		{"'Book' là?", []string{"Sách", "Bút", "Vở", "Bàn"}, 0},
		{"'Yes' nghĩa là gì?", []string{"Có", "Không", "Chắc chắn", "Không biết"}, 0},
		{"'How are you?' là?", []string{"Bạn khỏe không?", "Bạn tên gì?", "Bạn bao nhiêu tuổi?", "Bạn sống ở đâu?"}, 0},
		{"'Cat' nghĩa là gì?", []string{"Con mèo", "Con chó", "Con gà", "Con heo"}, 0},
		{"'Goodbye' là gì?", []string{"Tạm biệt", "Chào buổi sáng", "Xin chào", "Hẹn gặp lại"}, 0},
		{"'Red' là màu gì?", []string{"Đỏ", "Xanh", "Vàng", "Tím"}, 0},
	}
}

func genITQuestions() []RawQuestion {
	return []RawQuestion{
		{"HTML là viết tắt của gì?", []string{
			"HyperText Markup Language", "HighText Machine Language", "Hyper Transfer Markup Language", "Hyperlink and Text Markup Language"}, 0},
		{"Hệ điều hành phổ biến nhất hiện nay là?", []string{
			"Windows", "Linux", "macOS", "Ubuntu"}, 0},
		{"CPU là gì?", []string{
			"Central Processing Unit", "Computer Primary Unit", "Control Processing Unit", "Central Performance Unit"}, 0},
		{"Công cụ dùng để quản lý mã nguồn?", []string{
			"Git", "Photoshop", "Excel", "Word"}, 0},
		{"HTTP là giao thức gì?", []string{
			"Truyền siêu văn bản", "Bảo mật truyền tải", "Mã hóa dữ liệu", "Lưu trữ web"}, 0},
		{"IDE nào phổ biến cho lập trình Java?", []string{
			"IntelliJ IDEA", "Photoshop", "Premiere", "AutoCAD"}, 0},
		{"RAM dùng để làm gì?", []string{
			"Lưu dữ liệu tạm thời", "Lưu trữ dài hạn", "Kết nối mạng", "Chạy phần mềm diệt virus"}, 0},
		{"Đơn vị lưu trữ lớn nhất trong danh sách?", []string{
			"Terabyte", "Megabyte", "Kilobyte", "Gigabyte"}, 0},
		{"Ngôn ngữ phổ biến cho phát triển web?", []string{
			"JavaScript", "C#", "Swift", "Go"}, 0},
		{"Công ty phát triển hệ điều hành Android?", []string{
			"Google", "Microsoft", "Apple", "Samsung"}, 0},
	}
}

func genIQQuestions() []RawQuestion {
	return []RawQuestion{
		{"Số tiếp theo của dãy: 2, 4, 8, 16, ?", []string{"32", "24", "30", "36"}, 0},
		{"Nếu tất cả mèo là thú, và một số thú là chó, thì...", []string{"Một số mèo có thể là chó", "Tất cả chó là mèo", "Không có con chó nào là mèo", "Mèo không là gì cả"}, 2},
		{"Hình nào khác nhóm: hình tròn, hình vuông, hình tam giác, hình cầu?", []string{"Hình cầu", "Hình vuông", "Hình tròn", "Hình tam giác"}, 0},
		{"5 + 3 x 2 = ?", []string{"11", "16", "13", "10"}, 0},
		{"Nếu hôm qua là thứ Năm, thì ngày mai là?", []string{"Thứ Bảy", "Thứ Hai", "Chủ Nhật", "Thứ Sáu"}, 0},
		{"Từ nào khác nhóm: xe đạp, xe máy, tàu hỏa, máy bay?", []string{"Xe đạp", "Xe máy", "Tàu hỏa", "Máy bay"}, 0},
		{"Hình phản chiếu của số 3 là?", []string{"Ɛ", "E", "3", "W"}, 0},
		{"Nếu A=1, B=2... Z=26 thì G+O = ?", []string{"22", "17", "21", "15"}, 0},
		{"Chọn từ không cùng loại: đỏ, xanh, tròn, vàng?", []string{"Tròn", "Đỏ", "Xanh", "Vàng"}, 0},
		{"Con nào không giống: chó, mèo, cá, sách?", []string{"Sách", "Chó", "Mèo", "Cá"}, 0},
	}
}
func genLiteratureQuestions() []RawQuestion {
	return []RawQuestion{
		{"Tác giả 'Truyện Kiều' là ai?", []string{"Nguyễn Du", "Nguyễn Trãi", "Hồ Xuân Hương", "Tố Hữu"}, 0},
		{"Tác phẩm nào thuộc thể loại truyện ngắn?", []string{"Chiếc lược ngà", "Truyện Kiều", "Chinh phụ ngâm", "Lục Vân Tiên"}, 0},
		{"Câu thơ 'Bầu ơi thương lấy bí cùng' thuộc thể loại?", []string{"Ca dao", "Ngụ ngôn", "Tục ngữ", "Thơ hiện đại"}, 0},
		{"Tác phẩm 'Lão Hạc' do ai viết?", []string{"Nam Cao", "Ngô Tất Tố", "Nguyễn Huy Tưởng", "Tô Hoài"}, 0},
		{"Tác phẩm 'Dế Mèn phiêu lưu ký' là của?", []string{"Tô Hoài", "Xuân Quỳnh", "Nguyễn Nhật Ánh", "Nguyễn Du"}, 0},
		{"'Văn học dân gian' là?", []string{"Tác phẩm truyền miệng", "Tác phẩm viết", "Văn học hiện đại", "Văn học cổ điển châu Âu"}, 0},
		{"Từ 'nhân hóa' có nghĩa là?", []string{"Gán đặc điểm người cho vật", "Nói ngược", "Khen quá mức", "So sánh trực tiếp"}, 0},
		{"Thể thơ lục bát gồm mấy câu?", []string{"2", "3", "4", "1"}, 0},
		{"Truyện cười nhằm mục đích gì?", []string{"Giải trí và phê phán", "Tuyên truyền", "Giáo dục nghiêm túc", "Bình luận lịch sử"}, 0},
		{"Nhà văn nổi bật thời kỳ hiện đại?", []string{"Nguyễn Minh Châu", "Nguyễn Trãi", "Nguyễn Du", "Trần Quốc Tuấn"}, 0},
	}
}
func genMusicQuestions() []RawQuestion {
	return []RawQuestion{
		{"Nốt đầu tiên trong âm giai đô trưởng là?", []string{"Đô", "Rê", "Mi", "Fa"}, 0},
		{"Nhạc sĩ viết 'Tiến quân ca'?", []string{"Văn Cao", "Trịnh Công Sơn", "Phạm Duy", "Đoàn Chuẩn"}, 0},
		{"Nhạc cụ truyền thống Việt Nam?", []string{"Đàn tranh", "Guitar", "Piano", "Saxophone"}, 0},
		{"Một bản nhạc thường bắt đầu bằng?", []string{"Khóa nhạc", "Nốt cuối", "Quãng tám", "Tempo"}, 0},
		{"Ai là 'ông hoàng nhạc pop' thế giới?", []string{"Michael Jackson", "Elvis Presley", "Freddie Mercury", "Justin Bieber"}, 0},
		{"Nốt 'La' tương ứng tần số bao nhiêu?", []string{"440Hz", "220Hz", "880Hz", "320Hz"}, 0},
		{"Nhạc cụ nào dùng miệng để thổi?", []string{"Sáo", "Trống", "Đàn bầu", "Violon"}, 0},
		{"Phong cách 'acoustic' nghĩa là?", []string{"Dùng nhạc cụ mộc", "Điện tử hóa", "Không có nhạc", "Opera"}, 0},
		{"Nốt nhạc 'Sol' đứng sau nốt nào?", []string{"Fa", "Mi", "La", "Đô"}, 0},
		{"Người hát bài 'Bài ca không quên'?", []string{"Cẩm Vân", "Mỹ Tâm", "Sơn Tùng", "Đàm Vĩnh Hưng"}, 0},
	}
}

func genArtQuestions() []RawQuestion {
	return []RawQuestion{
		{"Chủ nghĩa ấn tượng bắt đầu ở nước nào?", []string{"Pháp", "Ý", "Tây Ban Nha", "Anh"}, 0},
		{"Danh họa vẽ 'Mona Lisa'?", []string{"Leonardo da Vinci", "Van Gogh", "Pablo Picasso", "Michelangelo"}, 0},
		{"Tượng Nữ thần Tự do là quà của quốc gia nào?", []string{"Pháp", "Anh", "Ý", "Canada"}, 0},
		{"Chất liệu 'tranh lụa' phổ biến ở?", []string{"Việt Nam", "Nhật", "Trung Quốc", "Pháp"}, 0},
		{"Mỹ thuật thời kỳ nào có Kim Tự Tháp?", []string{"Ai Cập cổ đại", "La Mã", "Hy Lạp", "Maya"}, 0},
		{"Tranh 'Hoa hướng dương' của ai?", []string{"Van Gogh", "Da Vinci", "Gauguin", "Raphael"}, 0},
		{"Tượng 'David' nổi tiếng của?", []string{"Michelangelo", "Donatello", "Bernini", "Rodin"}, 0},
		{"Hình học trong hội họa gọi là gì?", []string{"Hình khối", "Hình nền", "Chấm phá", "Tỷ lệ"}, 0},
		{"Màu nóng gồm?", []string{"Đỏ, cam, vàng", "Xanh, tím", "Đen, trắng", "Tím, hồng, xanh"}, 0},
		{"Tác phẩm điêu khắc thường dùng chất liệu nào?", []string{"Đá, gỗ, đồng", "Nhựa, giấy", "Vải, len", "Thủy tinh, nhôm"}, 0},
	}
}

func genSportQuestions() []RawQuestion {
	return []RawQuestion{
		{"Môn thể thao vua là?", []string{"Bóng đá", "Bóng chuyền", "Bóng rổ", "Bóng bàn"}, 0},
		{"Số người mỗi đội trong bóng đá?", []string{"11", "10", "9", "12"}, 0},
		{"Giải bóng đá lớn nhất hành tinh?", []string{"World Cup", "UEFA Champions League", "Premier League", "Copa America"}, 0},
		{"VĐV nổi tiếng Michael Jordan thi đấu môn?", []string{"Bóng rổ", "Bóng chày", "Bóng đá", "Quần vợt"}, 0},
		{"Môn thi đấu có vợt và cầu lông?", []string{"Cầu lông", "Bóng bàn", "Quần vợt", "Bóng rổ"}, 0},
		{"Thế vận hội còn gọi là gì?", []string{"Olympic", "SEA Games", "ASIAN Cup", "World Tour"}, 0},
		{"Quả bóng vàng 2022 thuộc về ai?", []string{"Lionel Messi", "Cristiano Ronaldo", "Benzema", "Mbappe"}, 0},
		{"Sân tennis có mấy loại mặt sân?", []string{"3", "2", "4", "5"}, 0},
		{"Môn thể thao đua xe nổi tiếng?", []string{"F1", "MotoGP", "Tour de France", "Rally"}, 0},
		{"VĐV bơi lội huyền thoại người Mỹ?", []string{"Michael Phelps", "Usain Bolt", "LeBron James", "Roger Federer"}, 0},
	}
}

func genEnvironmentQuestions() []RawQuestion {
	return []RawQuestion{
		{"Khí nào gây hiệu ứng nhà kính mạnh nhất?", []string{"CO2", "O2", "N2", "H2"}, 0},
		{"Tái chế là gì?", []string{"Sử dụng lại chất thải", "Vứt rác đi", "Đốt rác", "Chôn lấp rác"}, 0},
		{"Tài nguyên nào là tái tạo được?", []string{"Gió", "Than đá", "Dầu mỏ", "Khí thiên nhiên"}, 0},
		{"Rừng có vai trò gì?", []string{"Hấp thụ CO2", "Thải CO2", "Sản xuất bụi", "Gây hạn hán"}, 0},
		{"Ô nhiễm không khí gây ra bởi?", []string{"Khí thải xe", "Gió", "Mưa", "Ánh sáng"}, 0},
		{"Hiện tượng nóng lên toàn cầu gọi là?", []string{"Biến đổi khí hậu", "Mùa đông", "Sương mù", "Lốc xoáy"}, 0},
		{"Loài nào bị đe dọa do mất môi trường sống?", []string{"Hổ", "Chó", "Gà", "Heo"}, 0},
		{"Rác thải nhựa nên được?", []string{"Tái chế", "Đốt", "Vứt biển", "Chôn"}, 0},
		{"Nguồn năng lượng sạch?", []string{"Năng lượng mặt trời", "Than", "Dầu diesel", "Gas"}, 0},
		{"Môi trường ảnh hưởng trực tiếp đến?", []string{"Sức khỏe", "Thời trang", "Kinh tế", "Ẩm thực"}, 0},
	}
}

func genWorldCultureQuestions() []RawQuestion {
	return []RawQuestion{
		{"Lễ hội Halloween bắt nguồn từ đâu?", []string{"Châu Âu", "Châu Á", "Châu Mỹ", "Châu Phi"}, 0},
		{"Tết Âm lịch phổ biến ở quốc gia nào?", []string{"Việt Nam", "Anh", "Pháp", "Mỹ"}, 0},
		{"Kimono là trang phục truyền thống của?", []string{"Nhật Bản", "Hàn Quốc", "Trung Quốc", "Thái Lan"}, 0},
		{"Pizza có nguồn gốc từ nước nào?", []string{"Ý", "Pháp", "Mỹ", "Tây Ban Nha"}, 0},
		{"Điệu nhảy Tango bắt nguồn từ?", []string{"Argentina", "Brazil", "Tây Ban Nha", "Peru"}, 0},
		{"Tôn giáo chính ở Ấn Độ?", []string{"Hindu giáo", "Thiên chúa giáo", "Hồi giáo", "Phật giáo"}, 0},
		{"Ngôn ngữ phổ biến nhất thế giới?", []string{"Tiếng Anh", "Tiếng Hoa", "Tiếng Tây Ban Nha", "Tiếng Pháp"}, 0},
		{"Lễ hội té nước Songkran thuộc nước nào?", []string{"Thái Lan", "Lào", "Myanmar", "Indonesia"}, 0},
		{"Bia được uống nhiều ở nước nào?", []string{"Đức", "Pháp", "Ý", "Anh"}, 0},
		{"Văn hóa 'karaoke' phổ biến ở?", []string{"Nhật Bản", "Hàn Quốc", "Trung Quốc", "Việt Nam"}, 0},
	}
}

func genEthicsQuestions() []RawQuestion {
	return []RawQuestion{
		{"Đạo đức là gì?", []string{"Hành vi đúng đắn", "Luật pháp", "Tôn giáo", "Văn hóa"}, 0},
		{"Hành động nào thể hiện sự trung thực?", []string{"Nói thật", "Nói dối", "Giấu diếm", "Đổ lỗi"}, 0},
		{"Tôn trọng người khác là?", []string{"Lắng nghe & không xúc phạm", "Tranh cãi", "Bỏ qua", "Xem thường"}, 0},
		{"Lòng biết ơn thể hiện qua?", []string{"Cảm ơn, giúp đỡ lại", "Im lặng", "Phớt lờ", "Chê bai"}, 0},
		{"Ghen tị là hành vi?", []string{"Tiêu cực", "Tốt đẹp", "Trung lập", "Bình thường"}, 0},
		{"Giữ lời hứa giúp tăng?", []string{"Uy tín", "Tiền bạc", "Quyền lực", "Sự lười biếng"}, 0},
		{"Tha thứ là hành vi?", []string{"Khoan dung", "Báo thù", "Vô cảm", "Bướng bỉnh"}, 0},
		{"Đạo đức nghề nghiệp là?", []string{"Chuẩn mực hành vi trong công việc", "Sở thích cá nhân", "Luật cấm", "Khen thưởng"}, 0},
		{"Công bằng là khi?", []string{"Mọi người được đối xử như nhau", "Người giàu được ưu tiên", "Phân biệt đối xử", "Không rõ ràng"}, 0},
		{"Sống có trách nhiệm nghĩa là?", []string{"Dám làm dám chịu", "Trốn tránh", "Lười biếng", "Đùn đẩy"}, 0},
	}
}

func genLifeSkillQuestions() []RawQuestion {
	return []RawQuestion{
		{"Kỹ năng lắng nghe tốt là gì?", []string{"Chú ý và không ngắt lời", "Nói xen vào", "Nghe một nửa", "Lơ đãng"}, 0},
		{"Kỹ năng giải quyết xung đột là?", []string{"Thảo luận và tìm giải pháp", "Tránh mặt", "Nổi nóng", "Phớt lờ"}, 0},
		{"Kỹ năng quản lý thời gian giúp gì?", []string{"Hiệu quả công việc", "Ngủ nhiều", "Tăng stress", "Chậm tiến độ"}, 0},
		{"Làm việc nhóm hiệu quả cần?", []string{"Hợp tác và giao tiếp", "Im lặng", "Làm một mình", "Ra lệnh"}, 0},
		{"Kỹ năng ra quyết định tốt là?", []string{"Xem xét lựa chọn và hậu quả", "Chọn đại", "Phụ thuộc người khác", "Chờ đợi"}, 0},
		{"Tự tin là biểu hiện của?", []string{"Hiểu rõ bản thân", "Kiêu ngạo", "Chủ quan", "Mặc cảm"}, 0},
		{"Giữ bình tĩnh trong tình huống căng thẳng giúp?", []string{"Giải quyết tốt hơn", "Tăng mâu thuẫn", "Trốn tránh", "Bỏ cuộc"}, 0},
		{"Kỹ năng đặt mục tiêu giúp?", []string{"Có định hướng rõ ràng", "Làm việc tuỳ hứng", "Trì hoãn", "Làm gì cũng được"}, 0},
		{"Tự lập là gì?", []string{"Tự lo liệu và chịu trách nhiệm", "Phụ thuộc người khác", "Yêu cầu giúp đỡ", "Phó mặc"}, 0},
		{"Giao tiếp hiệu quả cần?", []string{"Ngắn gọn, rõ ràng, lịch sự", "Nói dài", "Gắt gỏng", "Không nói"}, 0},
	}
}
func genSocialScienceQuestions() []RawQuestion {
	return []RawQuestion{
		{"Xã hội học nghiên cứu về?", []string{"Con người và xã hội", "Thực vật", "Hóa chất", "Cơ khí"}, 0},
		{"Nhà nước là?", []string{"Tổ chức quyền lực chính trị", "Câu lạc bộ", "Công ty", "Gia đình"}, 0},
		{"Pháp luật nhằm?", []string{"Điều chỉnh hành vi xã hội", "Tuyên truyền văn hóa", "Khai thác tài nguyên", "Giải trí"}, 0},
		{"Gia đình là?", []string{"Tế bào của xã hội", "Tổ chức chính trị", "Xã hội thu nhỏ", "Đơn vị pháp lý"}, 0},
		{"Văn hóa là?", []string{"Tổng thể giá trị vật chất và tinh thần", "Nghệ thuật", "Giải trí", "Thể thao"}, 0},
		{"Giai cấp là?", []string{"Nhóm người có địa vị kinh tế giống nhau", "Tập thể", "Tôn giáo", "Giới tính"}, 0},
		{"Quyền con người là?", []string{"Quyền cơ bản của mỗi người", "Đặc quyền nhà nước", "Cấp phát", "Chỉ áp dụng cho người lớn"}, 0},
		{"Dân chủ là?", []string{"Quyền lực thuộc về nhân dân", "Chính quyền độc tài", "Quân chủ", "Chỉ dành cho vua"}, 0},
		{"Tôn giáo có vai trò gì?", []string{"Hướng dẫn đạo đức và niềm tin", "Kinh tế", "Giải trí", "Thời trang"}, 0},
		{"Cộng đồng là?", []string{"Nhóm người chung lợi ích", "Doanh nghiệp", "Cá nhân", "Tổ đội"}, 0},
	}
}
func genVietnameseQuestions() []RawQuestion {
	return []RawQuestion{
		{"Danh từ là gì?", []string{"Từ chỉ người, vật, sự việc", "Từ chỉ hành động", "Từ miêu tả", "Từ nối"}, 0},
		{"Câu ghép là gì?", []string{"Câu có 2 vế có quan hệ với nhau", "Câu có một chủ ngữ", "Câu cảm thán", "Câu nghi vấn"}, 0},
		{"Từ đồng nghĩa là?", []string{"Từ khác nhau nhưng nghĩa giống", "Từ có nghĩa trái nhau", "Từ nhiều nghĩa", "Từ mượn"}, 0},
		{"Từ láy là?", []string{"Từ có âm lặp lại", "Từ chỉ sự việc", "Từ mượn", "Từ ghép"}, 0},
		{"Câu rút gọn là?", []string{"Câu lược bỏ thành phần", "Câu dài", "Câu hỏi", "Câu đảo ngữ"}, 0},
		{"Trạng ngữ chỉ?", []string{"Thời gian, nơi chốn, cách thức", "Chủ ngữ", "Động từ", "Tân ngữ"}, 0},
		{"Từ trái nghĩa là?", []string{"Từ có nghĩa ngược nhau", "Từ mượn", "Từ đa nghĩa", "Từ đồng âm"}, 0},
		{"Phó từ là gì?", []string{"Từ bổ sung ý nghĩa cho động/tính từ", "Từ chỉ người", "Từ nối", "Từ cảm thán"}, 0},
		{"Từ tượng thanh mô phỏng?", []string{"Âm thanh", "Màu sắc", "Cảm xúc", "Kích thước"}, 0},
		{"Từ nghi vấn dùng để?", []string{"Hỏi", "Miêu tả", "Cảm thán", "Trình bày"}, 0},
	}
}
func genCodingQuestions() []RawQuestion {
	return []RawQuestion{
		{"Ngôn ngữ nào phổ biến để viết web frontend?", []string{"JavaScript", "Python", "Go", "Java"}, 0},
		{"Biến trong lập trình dùng để?", []string{"Lưu trữ dữ liệu", "In ra màn hình", "Vẽ hình", "Tăng tốc CPU"}, 0},
		{"Cấu trúc điều kiện thường dùng?", []string{"if-else", "print", "input", "loop"}, 0},
		{"Vòng lặp dùng để?", []string{"Lặp lại đoạn mã", "Thoát chương trình", "Chạy 1 lần", "Gán giá trị"}, 0},
		{"Ngôn ngữ nào được biên dịch?", []string{"C/C++", "HTML", "CSS", "SQL"}, 0},
		{"Mảng (array) là gì?", []string{"Tập hợp nhiều giá trị", "Câu lệnh điều kiện", "Thư viện ảnh", "Biến boolean"}, 0},
		{"Lập trình hướng đối tượng viết tắt là?", []string{"OOP", "POP", "PLO", "LOL"}, 0},
		{"IDE là?", []string{"Môi trường phát triển tích hợp", "Trình duyệt", "Cơ sở dữ liệu", "Máy chủ"}, 0},
		{"Go được phát triển bởi?", []string{"Google", "Apple", "Microsoft", "IBM"}, 0},
		{"Git dùng để?", []string{"Quản lý mã nguồn", "Chạy server", "Tạo slide", "Thiết kế UI"}, 0},
	}
}
func genScienceQuestions() []RawQuestion {
	return []RawQuestion{
		{"Nước sôi ở bao nhiêu độ C?", []string{"100", "90", "80", "120"}, 0},
		{"Cơ quan hô hấp của người?", []string{"Phổi", "Gan", "Tim", "Dạ dày"}, 0},
		{"Tốc độ ánh sáng gần đúng?", []string{"300,000 km/s", "150,000 km/s", "100,000 km/s", "1 triệu km/s"}, 0},
		{"Mắt người nhìn thấy ánh sáng nào?", []string{"Ánh sáng khả kiến", "Hồng ngoại", "Tia X", "Tia gamma"}, 0},
		{"Lực hút của Trái Đất gọi là gì?", []string{"Trọng lực", "Ma sát", "Từ lực", "Động lực"}, 0},
		{"Máu đỏ do gì tạo ra?", []string{"Hồng cầu", "Bạch cầu", "Tiểu cầu", "Tủy xương"}, 0},
		{"Thực vật cần gì để quang hợp?", []string{"Ánh sáng", "Oxy", "Nhiệt độ", "Gió"}, 0},
		{"Âm thanh lan truyền qua?", []string{"Không khí", "Chân không", "Ánh sáng", "Điện"}, 0},
		{"Hệ mặt trời có mấy hành tinh?", []string{"8", "9", "7", "10"}, 0},
		{"Trái đất quay quanh gì?", []string{"Mặt trời", "Mặt trăng", "Sao Hỏa", "Trục"}, 0},
	}
}

func genMathQuestions() []RawQuestion {
	return []RawQuestion{
		{"1 + 1 = ?", []string{"2", "3", "1", "0"}, 0},
		{"5 × 6 = ?", []string{"30", "11", "56", "26"}, 0},
		{"10 ÷ 2 = ?", []string{"5", "2", "8", "10"}, 0},
		{"Căn bậc hai của 25 là?", []string{"5", "4", "6", "3"}, 0},
		{"20% của 50 là?", []string{"10", "5", "15", "20"}, 0},
		{"Số nguyên tố đầu tiên là?", []string{"2", "3", "1", "5"}, 0},
		{"10 + 15 = ?", []string{"25", "20", "15", "30"}, 0},
		{"50 - 18 = ?", []string{"32", "42", "28", "38"}, 0},
		{"9 × 9 = ?", []string{"81", "72", "91", "99"}, 0},
		{"7² = ?", []string{"49", "36", "64", "25"}, 0},
	}
}
