package content

import (
	"testing"
)

func TestExtract(t *testing.T) {

	content := `​​虎扑5月28日讯 在当前世界杯即将打响之际，众多已退役的大牌球星纷纷赴中国参加商业活动捞金，据澎湃新闻网报道，蚌埠召开的“世界足球峰会”，光请贝克汉姆一人的费用就超过了人民币2000万元。

	据统计，俄罗斯世界杯前一周，国内各地跟世界杯相关的活动就有数十场，邀请到的各路传奇巨星也着实不少。日前，安徽蚌埠的“世界足球峰会”还请到了“万人迷”贝克汉姆。当然，小贝驾临蚌埠这样的四线城市肯定是天价。
	
	一位知情人表示，“具体价格不能透露，但可以负责任的说，小贝一个人足以顶上20个普约尔。”而据了解，前后各种花销加起来，请小贝的总开支超过了2000万元人民币。
	
	随着中超最近的火热，越来越多的大牌球星已经不介意来到中国赚快钱。而对于像欧文、皮耶罗、普约尔、古利特这样还没来得及体会中超的“金元风暴”就已经退役的球星们来说，就只能通过商业活动来捞金了。这几年时间，不但中超联赛大牌外援薪水飞涨，就连一些退役大牌明星来到中国出席活动的出场费也跟着提高。
	
	一位来自上海的中国商务经纪人提到了欧文，“欧文这两个月来了多少趟中国了，这么多企业给钱，他们能不开心吗？”而即便像古利特这样年纪更大的老球星也不乏市场，4月上海，5月北京，6月蚌埠，一样可以赚得盆满钵满。“古利特的一个中国朋友还找过我，说特别想来中国执教，在欧洲，他们哪里能赚这些钱。”
	
	中国的世界杯热潮在6月2日这天体现得淋漓尽致，一个上海就有两场与欧美球星有关的大型活动。某家门户网站举办的名人足球赛，请来了普约尔、萨内蒂、卡洛斯等不少明星，很多承诺来参加比赛的明星要价都翻了番。同时，上海市足协还搭台举办了国际传奇球星邀请赛，邀请了埃芬博格、卡普德维拉、马尼切、西芒等众多欧美足坛元老参赛。
	
	而除此之外，像罗纳尔多、卡洛斯这样的高知名度球星，也在纷纷进行自己的中国行活动。
	
	“这些人来中国肯定就是淘金的，不趁现在涨价行嘛！”一位赛事组织者介绍，像内斯塔，普约尔，吉格斯这些退役大牌，在世界杯前身价至少翻了一倍。“现在他们基本上10万欧是起步价。”一位经常运作欧美球星来华的s姓经纪人说。
	
	不过，这些球星的身价，在“万人迷”贝克汉姆的天价面前也只能算小巫。6月6日，贝克汉姆将空降安徽蚌埠，参加“世界足球峰会”，据了解，主办方加上包机等，付出的代价超过2000万人民币。这让一位曾经与小贝团队合作多次的某跨国体育公司高管也不禁感叹，“中国不缺钱，更不缺想象力。”
	
	“小贝本来就是足坛第一身价，肯定是让普约尔们难以望其项背的，何况他最近已经不再愿意参加普通的活动论坛。这次蚌埠方面能请动他肯定是花了血本的，之前我也请过他参加一个企业的活动，但没想到他会去蚌埠。”`

	keywords := Extract(content, 10)
	t.Logf("keywords: %v", keywords)

	content = `​​北京时间31日早间消息，彭博社援引不具名知情人士的话称，美国亿万富翁沃伦-巴菲特今年较早时曾提出向优步科技公司注资30亿美元，但由于双方在交易条款上存在分歧，谈判最终失败了。

	　　该报道称，巴菲特旗下伯克希尔哈撒韦(292000, 5999.99, 2.10%)公司提出向优步提供一笔可转换贷款，如果这家硅谷叫车公司遭遇财政危机，这将为巴菲特的投资提供保护。
	
	　　但优步CEO达拉-科斯罗萨西（Dara Khosrowshahi）提议将这笔交易的规模缩减至20亿美元，让巴菲特在该公司持有较小的股份。由于双方未能就交易条款达成一致，谈判最终告吹。
	
	　　巴菲特在接受CNBC采访时称，彭博社的报道“部分细节不正确”，但他证实，伯克希尔的确曾与优步进行谈判。
	
	　　巴菲特过去一直避免大幅投资科技公司，但近年来投资策略有所转变，他旗下的伯克希尔从2016年开始投资苹果(187.5, -0.40, -0.21%)公司，目前已成为该公司第二大股东。此外，巴菲特还在多个场合对没有在谷歌(1067.8, 7.48, 0.71%)和亚马逊(1624.89, 12.02, 0.75%)做大之前进行投资表示遗憾。`

	keywords = Extract(content, 10)
	t.Logf("keywords: %v", keywords)
}

func BenchmarkExtract(b *testing.B) {
	content := `​​虎扑5月28日讯 在当前世界杯即将打响之际，众多已退役的大牌球星纷纷赴中国参加商业活动捞金，据澎湃新闻网报道，蚌埠召开的“世界足球峰会”，光请贝克汉姆一人的费用就超过了人民币2000万元。

	据统计，俄罗斯世界杯前一周，国内各地跟世界杯相关的活动就有数十场，邀请到的各路传奇巨星也着实不少。日前，安徽蚌埠的“世界足球峰会”还请到了“万人迷”贝克汉姆。当然，小贝驾临蚌埠这样的四线城市肯定是天价。
	
	一位知情人表示，“具体价格不能透露，但可以负责任的说，小贝一个人足以顶上20个普约尔。”而据了解，前后各种花销加起来，请小贝的总开支超过了2000万元人民币。
	
	随着中超最近的火热，越来越多的大牌球星已经不介意来到中国赚快钱。而对于像欧文、皮耶罗、普约尔、古利特这样还没来得及体会中超的“金元风暴”就已经退役的球星们来说，就只能通过商业活动来捞金了。这几年时间，不但中超联赛大牌外援薪水飞涨，就连一些退役大牌明星来到中国出席活动的出场费也跟着提高。
	
	一位来自上海的中国商务经纪人提到了欧文，“欧文这两个月来了多少趟中国了，这么多企业给钱，他们能不开心吗？”而即便像古利特这样年纪更大的老球星也不乏市场，4月上海，5月北京，6月蚌埠，一样可以赚得盆满钵满。“古利特的一个中国朋友还找过我，说特别想来中国执教，在欧洲，他们哪里能赚这些钱。”
	
	中国的世界杯热潮在6月2日这天体现得淋漓尽致，一个上海就有两场与欧美球星有关的大型活动。某家门户网站举办的名人足球赛，请来了普约尔、萨内蒂、卡洛斯等不少明星，很多承诺来参加比赛的明星要价都翻了番。同时，上海市足协还搭台举办了国际传奇球星邀请赛，邀请了埃芬博格、卡普德维拉、马尼切、西芒等众多欧美足坛元老参赛。
	
	而除此之外，像罗纳尔多、卡洛斯这样的高知名度球星，也在纷纷进行自己的中国行活动。
	
	“这些人来中国肯定就是淘金的，不趁现在涨价行嘛！”一位赛事组织者介绍，像内斯塔，普约尔，吉格斯这些退役大牌，在世界杯前身价至少翻了一倍。“现在他们基本上10万欧是起步价。”一位经常运作欧美球星来华的s姓经纪人说。
	
	不过，这些球星的身价，在“万人迷”贝克汉姆的天价面前也只能算小巫。6月6日，贝克汉姆将空降安徽蚌埠，参加“世界足球峰会”，据了解，主办方加上包机等，付出的代价超过2000万人民币。这让一位曾经与小贝团队合作多次的某跨国体育公司高管也不禁感叹，“中国不缺钱，更不缺想象力。”
	
	“小贝本来就是足坛第一身价，肯定是让普约尔们难以望其项背的，何况他最近已经不再愿意参加普通的活动论坛。这次蚌埠方面能请动他肯定是花了血本的，之前我也请过他参加一个企业的活动，但没想到他会去蚌埠。”`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Extract(content, 6)
	}
}
