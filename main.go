package main

import "fmt"

func CardReader() chan string {
	cr := make(chan string)

	data := [...]string{
		`1234567890qwertyuiop[]asdfghjkl;'zxc**vbnm,./?><":}{+_)(*&^%$#@!ASDfadsgfaFGGfat`,
		`jaskd;lkjL:LKJHA;lkjdsaf;lkasdf;lkhf'L:KAS:FJBADF:"Lkhasd;flkjjasdfl;KHASD;lhlka`,
		`J:AKSJD:Jba;slkjhdf;jabsd;flkHA:dfkba'dflkH"DFlkh"LASDkhfashfoaweioi**A(YASPjkln`,
		`:lksdnf;jkalnadsfpubnqpjknap;fvigohaewiulenf;lkdfnv;aksfdyg;oiat7ahjasmdfn;lsdfk`,
		`;aslkdnf89h4a;lkjah854;lnkaf845;lrdkan49aerlprkgn95;lrdaegknq9o5;oqalekgo598h;sl`,
		`aesdfja;sdlkhf;loashd;fljkhas;ldfka;sdlkfh;asdklhjf;asdkljf;lasdkjf;lsdjkaf;lksj`,
		`adf;lksjad;flkanviuerhp34hioef980q243rp93nfp9*H(*AH(&HW(*#Rghp3nu4fpq9834ghfp9*H`,
		`W$(8hgqp3a4hfu90pa83gfpOIUYHSEdapfoiahp0f;askldjf;lsdjkaf;kajsd;flkasdfas'd;fias`,
		`dfa)sdfasdfasd)fasdfasdfasdfasdgfa)kasdlfkjhalsdjkhfljkashdlfkjhlkjhlkjhkhkjhasl`,
		`dkfjhaosuibghvlkjabsfdlvijuhbalsdkjfhflksdjabvlikjuaebolvibas<dfjvbalduskvfhblas`,
		`djkfhbvlkujashfo8iuqghiujashrdfv87tIPUATD*^TRAOCUYTf97^DEfgoIUGXC098AGco8uY^ASFD`,
		`(C6ga0ds89c6fgaowuigfcboiuGO*^AFSIUHUGA*SUDYFG)*#wted8a8wegd087qaw4tgfoUYAG )*#G`,
		`FDOUWC#G)*&@GDPIOuyg@#)*&DG#@*)&XG)#@*UDGO*#uydgoawsyuidgfco8aywegdflihagwsd9f87`,
		`6g(A*&^DTF(&*#@^TRE(&#@^TD(&*^T#&^T@#D(&^T#@D(*&^@#TdkjHASDGc087awetgdfliuasgdf0`,
		`87agsdfljkhagsdokfjghasiudyhfiasdhfiahsifhaksfhkashkfhasikufhg3wo84faioiuwfduiaf`,
		`weuioawefiuadsfuioh897tyliukwjegfh87q34gfoi34uqgfc87q34gfoiugdffco87q24r87q3y487`,
		`q4ytr87qy4tro4iujhgc8&AWEfgljkhdeswgafoiug0d89&TA*I&O^RT&^#YRKHJGEASOd8hgLKJHGD*`,
		`&#EGWQDO*#&@TGOI#WUGD#@*UIG#W)*OD&G@#*O&#@GOIUTG@OIugaewoiuagfsdfioutaosduiyfiou`,
		`aysdifyisyfisyfisyf78a4ylfkjjqw430f789ytwaoeuidjfgoasuiydgfcio87awsgdflijuahgw08`,
		`74y3eprft3uqyr4f0**fsghr986y10q4r98ftladsjkfc-a0wseduf['sdpfau[sd'fasdfu][a'sdfa`,
		`[sd'fas]d[f'aswd]'afsd[f'oajsdpofihASUIDghoqfhi89uqgpiuh1poihlkjh4rpouih13puihsd`,
		`f;gkja;skfdghpaosifdhg;lksdjaf;oisajdf;'lksadn;flksajd[fgoiAKJSGDIUGO*(&#TOUGOIU`,
		`G#O*&TG#@ORYUGO$*&RTGKJASDhflkjashdfoiuagsdoiftgao8s7dtf*&GliGAOc87tIUTIUGLJKHg8`,
		`7asdtf8JHBo87tWE**UYAT(&D^RTAOHG*A^%DFIAWTEDV%^$QAWDF(O&WQGXCP(*gpiuASghd08a7wet`,
		`0f87awstefoiuagsdflkjashdpfiohasdpioufyhasidhfpiosuahdfpiouahsdfkl;jhasdlckjhasd`,
		`oifugasdlkfjhasdlk;jfhasldjkfhlasdjkhflkstjadhflkjasdhflkjashdflkjahsdflkjahsdlk`,
		`fjhaw89ufrhgq3p9847fgh98q73p4y897q34yfpiju3hrco89iuhqf87ihy43qpfhurwef8p97h34987`,
		`hfoijushdfvlkjasdho98uihg487fliwuaehflkjdsahclfksjadhfo897q3h4go8fq97y3r087ty34t`,
		`oiugqhklfjahds98cv7qglrjkhasdwpof;kjhapw9iusd8fyajkdshfpoiashdfo;kjhklksdjhf;lks`,
		`dahfu9h4rtokhtiusikhtpoiu;h.asdfhniuasd;klfjhasw8uflksdjahfoikjht89i45yh98tr4y89`,
		`484y48ytqo9hdfkljahsd89hawefiohadsp98yhalsjkdhcvpa98dwyfilajksdhflksdjahflsdajkh`,
		`flasdjkhfahs;fgkjhnaspfuhq3wp[gjknqwep[0frivjas[dfg')asdfaskmdn.aksnd;lkans;dlkf`,
		`ja;sdlkjf;lsjkadfp0iha;rnjgqp938htg9i34ht9834y8943ypihLJH9i8yLijasdfasdf356djh78`}

	go func() {
		for _, card := range data {
			cr <- card
		}
		close(cr)
	}()

	return cr
}

func LinePrinter(feed chan string) {

	for card, ok := <-feed; ok; card, ok = <-feed {
		fmt.Println(card)
	}

}

func main() {
	cr := CardReader()

	feed := Assemble(Copy(Disassemble(cr)))

	LinePrinter(feed)
}
