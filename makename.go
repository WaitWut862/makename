package makename

import (
	"math/rand"
)

func CompleteRand() string {
	chance := rand.Intn(10)
	if chance < 6 {
		return FullNeutral()
	} else if chance == 6 || chance == 7 {
		return FirstLastNeutral()
	} else if chance == 8 {
		return FullMidInitialNeutral()
	} else if chance == 9 {
		return FirstLastInitialNeutral()
	}
	return ""
}

func FullNeutral() string {
	gender := rand.Intn(2)
	if gender == 0 {
		fn := maleFirstNames[rand.Intn(len(maleFirstNames))]
		mn := maleMiddleNames[rand.Intn(len(maleMiddleNames))]
		ln := lastNames[rand.Intn(len(lastNames))]
		fullName := fn + " " + mn + " " + ln
		return fullName
	} else {
		fn := femaleFirstNames[rand.Intn(len(femaleFirstNames))]
		mn := femaleMiddleNames[rand.Intn(len(femaleMiddleNames))]
		ln := lastNames[rand.Intn(len(lastNames))]
		fullName := fn + " " + mn + " " + ln
		return fullName
	}
}

func FullMale() string {
	fn := maleFirstNames[rand.Intn(len(maleFirstNames))]
	mn := maleMiddleNames[rand.Intn(len(maleMiddleNames))]
	ln := lastNames[rand.Intn(len(lastNames))]
	fullName := fn + " " + mn + " " + ln
	return fullName
}

func FullFemale() string {
	fn := femaleFirstNames[rand.Intn(len(femaleFirstNames))]
	mn := femaleMiddleNames[rand.Intn(len(femaleMiddleNames))]
	ln := lastNames[rand.Intn(len(lastNames))]
	fullName := fn + " " + mn + " " + ln
	return fullName
}

func FirstLastNeutral() string {
	gender := rand.Intn(2)
	if gender == 0 {
		fn := maleFirstNames[rand.Intn(len(maleFirstNames))]
		ln := lastNames[rand.Intn(len(lastNames))]
		fullName := fn + " " + ln
		return fullName
	} else {
		fn := femaleFirstNames[rand.Intn(len(femaleFirstNames))]
		ln := lastNames[rand.Intn(len(lastNames))]
		fullName := fn + " " + ln
		return fullName
	}
}

func FirstLastMale() string {
	fn := maleFirstNames[rand.Intn(len(maleFirstNames))]
	ln := lastNames[rand.Intn(len(lastNames))]
	fullName := fn + " " + ln
	return fullName
}

func FirstLastFemale() string {
	fn := femaleFirstNames[rand.Intn(len(femaleFirstNames))]
	ln := lastNames[rand.Intn(len(lastNames))]
	fullName := fn + " " + ln
	return fullName
}

func FirstNeutral() string {
	gender := rand.Intn(2)
	if gender == 0 {
		fn := maleFirstNames[rand.Intn(len(maleFirstNames))]
		return fn
	} else {
		fn := femaleFirstNames[rand.Intn(len(femaleFirstNames))]
		return fn
	}
}

func FirstMale() string {
	fn := maleFirstNames[rand.Intn(len(maleFirstNames))]
	return fn
}

func FirstFemale() string {
	fn := femaleFirstNames[rand.Intn(len(femaleFirstNames))]
	return fn
}

func LastName() string {
	ln := lastNames[rand.Intn(len(lastNames))]
	return ln
}

func FullMidInitialNeutral() string {
	gender := rand.Intn(2)
	if gender == 0 {
		fn := maleFirstNames[rand.Intn(len(maleFirstNames))]
		mi := initials[rand.Intn(len(initials))]
		ln := lastNames[rand.Intn(len(lastNames))]
		fullName := fn + " " + mi + " " + ln
		return fullName
	} else {
		fn := femaleFirstNames[rand.Intn(len(femaleFirstNames))]
		mi := initials[rand.Intn(len(initials))]
		ln := lastNames[rand.Intn(len(lastNames))]
		fullName := fn + " " + mi + " " + ln
		return fullName
	}
}

func FullMidInitialMale() string {
	fn := maleFirstNames[rand.Intn(len(maleFirstNames))]
	mi := initials[rand.Intn(len(initials))]
	ln := lastNames[rand.Intn(len(lastNames))]
	fullName := fn + " " + mi + " " + ln
	return fullName
}

func FullMidInitialFemale() string {
	fn := femaleFirstNames[rand.Intn(len(femaleFirstNames))]
	mi := initials[rand.Intn(len(initials))]
	ln := lastNames[rand.Intn(len(lastNames))]
	fullName := fn + " " + mi + " " + ln
	return fullName
}

func FirstLastInitialNeutral() string {
	gender := rand.Intn(2)
	if gender == 0 {
		fn := maleFirstNames[rand.Intn(len(maleFirstNames))]
		li := initials[rand.Intn(len(initials))]
		fullName := fn + " " + li
		return fullName
	} else {
		fn := femaleFirstNames[rand.Intn(len(femaleFirstNames))]
		li := initials[rand.Intn(len(initials))]
		fullName := fn + " " + li
		return fullName
	}
}

var femaleFirstNames = []string{
	"Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan", "Jessica", "Sarah", "Karen",
	"Nancy", "Margaret", "Lisa", "Betty", "Dorothy", "Sandra", "Ashley", "Kimberly", "Donna", "Emily",
	"Carol", "Michelle", "Amanda", "Melissa", "Deborah", "Stephanie", "Rebecca", "Laura", "Sharon", "Cynthia",
	"Kathleen", "Amy", "Shirley", "Angela", "Helen", "Anna", "Brenda", "Pamela", "Nicole", "Emma",
	"Samantha", "Katherine", "Christine", "Debra", "Rachel", "Catherine", "Carolyn", "Janet", "Ruth", "Maria",
	"Chloe", "Isabella", "Sofia", "Ava", "Mia", "Charlotte", "Abigail", "Harper", "Lily", "Ella",
	"Hannah", "Aria", "Scarlett", "Zoe", "Nora", "Layla", "Lillian", "Audrey", "Penelope", "Claire",
	"Lucy", "Stella", "Violet", "Savannah", "Brooklyn", "Bella", "Skylar", "Paisley", "Aaliyah", "Ellie",
	"Madelyn", "Genesis", "Maya", "Naomi", "Alice", "Ariana", "Leah", "Caroline", "Hazel", "Anna",
	"Serenity", "Autumn", "Eva", "Everly", "Hailey", "Natalie", "Aurora", "Sadie", "Isla", "Kennedy",
	"Addison", "Kaylee", "Peyton", "Sophie", "Gabriella", "Mackenzie", "Savannah", "Zoey", "Sienna",
	"Elena", "Emilia", "Vivian", "Faith", "Jasmine", "Lila", "Carla", "Valentina", "Isabelle", "Maya",
	"Aubrey", "Hailey", "Piper", "Jocelyn", "Madeline", "Adeline", "Stella", "Delilah", "Avery", "Olivia",
	"Autumn", "Ella", "Hazel", "Aurora", "Skyler", "Zara", "Leilani", "Lydia", "Sophie", "Reagan",
}

var maleFirstNames = []string{
	"James", "John", "Robert", "Michael", "William", "David", "Richard", "Joseph", "Thomas", "Charles",
	"Christopher", "Daniel", "Matthew", "Anthony", "Donald", "Mark", "Paul", "Steven", "Andrew", "Kenneth",
	"George", "Joshua", "Kevin", "Brian", "Edward", "Ronald", "Timothy", "Jason", "Jeffrey", "Ryan",
	"Jacob", "Gary", "Nicholas", "Eric", "Jonathan", "Stephen", "Larry", "Justin", "Scott", "Brandon",
	"Benjamin", "Samuel", "Gregory", "Frank", "Alexander", "Raymond", "Patrick", "Jack", "Dennis", "Jerry",
	"Caleb", "Lucas", "Henry", "Owen", "Leo", "Isaac", "Nathan", "Elijah", "Liam", "Jackson",
	"Sebastian", "Aiden", "Grayson", "Julian", "Levi", "Mateo", "Wyatt", "Gabriel", "Asher", "Lincoln",
	"Ezra", "Christian", "Thomas", "Hudson", "Cameron", "Hunter", "Eli", "Aaron", "Adrian", "Cooper",
	"Easton", "Miles", "Robert", "Jameson", "Adam", "Ian", "Carson", "Axel", "Maverick", "Declan",
	"Xavier", "Silas", "Nathaniel", "Brody", "Zachary", "Wesley", "Emmett", "Kingston", "Theo", "Bentley",
	"Matthew", "Aiden", "Jack", "Leo", "Mason", "David", "Luke", "Tyler", "Zane", "Carter",
	"Cooper", "Blake", "Tucker", "Jaxon", "Connor", "Miles", "Owen", "Avery", "William", "Elliot",
	"Dean", "Landon", "Chase", "Mason", "Bryce", "Kyle", "Colton", "Brock", "Weston", "Jace",
	"Dominic", "Jude", "Grayson", "Blaine", "Jackson", "Gage", "Zander", "Liam", "Charlie", "Wilder",
	"Zach", "Nolan", "Isaiah", "Gavin", "Tristan", "Tanner", "Graham", "Tate", "Rowan", "Brayden",
	"Dean", "Hunter", "Julian", "Bennett", "Weston", "Wyatt", "Jameson", "Maverick", "Brooks", "Oliver",
	"Evan", "Brody", "Xander", "Maddox", "Caden", "Elijah", "Parker", "Warren", "Axel", "Jasper",
}

var femaleMiddleNames = []string{
	"Marie", "Lynn", "Grace", "Rae", "Ann", "Jean", "Jade", "Claire", "Jane", "Hope",
	"Kay", "Skye", "Faith", "Reese", "Joy", "Rose", "Eve", "June", "Mae", "True",
	"Elle", "Brooke", "Shannon", "Renee", "Paige", "Nicole", "Lee", "Gail", "Terry", "Catherine",
	"Elizabeth", "Louise", "Brittany", "Samantha", "Caroline", "Helen", "Diane", "Barbara", "Rosemary", "Vera",
	"Michelle", "Anne", "Ellen", "Frances", "Joan", "Lynn", "Patricia", "Kristine", "Megan", "Valerie",
	"Christine", "Tina", "Alison", "Katherine", "Amanda", "Melanie", "Debra", "Victoria", "Monica", "Shirley",
	"Joanne", "Heather", "Linda", "Cynthia", "Dawn", "Jill", "Natalie", "Pamela", "Kelsey", "Tiffany",
	"Paula", "Suzanne", "Danielle", "Carla", "Lori", "Veronica", "Jessica", "Angela", "Diana", "Marcia",
	"Linda", "Cheryl", "Barbara", "Lynn", "Betty", "Teresa", "Kayla", "Courtney", "Sabrina", "Gina",
	"Tracy", "Martha", "Bonnie", "Kelly", "Patty", "Stephanie", "Robin", "Carrie", "Jody", "Joanne",
	"Gloria", "Nina", "Rita", "Adele", "Krista", "Edith", "Charlotte", "Beatrice", "Jasmine", "Tasha",
	"Faith", "Maggie", "Shelley", "Marissa", "April", "Brenda", "Georgina", "Cheryl", "Nancy", "Sarah",
	"Audrey", "Ellen", "Doris", "Michele", "Monica", "Sally", "Tessa", "Rosie", "Hannah", "Lola",
}

var maleMiddleNames = []string{
	"James", "John", "Ray", "Dale", "Scott", "Paul", "Dean", "Eli", "Blake", "Noel",
	"Quinn", "Drew", "Neil", "Wade", "Max", "Mark", "Shane", "Cole", "Glenn", "Tate",
	"Rey", "Luke", "Lane", "Joel", "Pierce", "Knox", "Ty", "Hayes", "Mack", "Jett",
	"Reid", "Zane", "Pax", "Flynn", "Troy", "Cade", "Sean", "Owen", "Brady", "Clark",
	"Ezra", "Beau", "Micah", "Rhys", "Logan", "Luca", "Asa", "Toby", "Charles", "Allen",
	"Samuel", "Joseph", "Matthew", "Henry", "Robert", "David", "Andrew", "George", "Edward", "Thomas",
	"William", "Michael", "Richard", "Stephen", "Kenneth", "Brian", "Frank", "Scott", "Donald", "Gregory",
	"Raymond", "Lucas", "Jack", "John", "Benjamin", "Paul", "Christian", "Jonathan", "Victor", "Elliot",
	"Daniel", "Dylan", "Cameron", "Jordan", "Anthony", "Chris", "Brian", "Philip", "Bennett", "George",
	"Lawrence", "Derrick", "Aaron", "Timothy", "Caleb", "Hunter", "Evan", "Mason", "Eric", "Justin",
	"Matthew", "Keith", "Francis", "Vince", "Maxwell", "Riley", "Brock", "Lyle", "Jesse", "Kurtis",
	"Russell", "Tanner", "Travis", "Shawn", "Tyler", "Xander", "Felix", "Bradley", "Omar", "Graham",
	"Steven", "Hugh", "Clayton", "Todd", "Damon", "Mitchell", "Reagan", "Jordan", "Victor", "Adam",
	"Tyson", "Arthur", "Marvin", "Trenton", "Harrison", "Brayden", "Greg", "Russell", "Landon", "Dwayne",
	"Nelson", "Dante", "Jared", "Quincy", "Elijah", "Dustin", "Murray", "Wesley", "Solomon", "Jared",
}

var lastNames = []string{
	"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
	"Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin",
	"Lee", "Perez", "Thompson", "White", "Harris", "Sanchez", "Clark", "Ramirez", "Lewis", "Robinson",
	"Walker", "Young", "Allen", "King", "Wright", "Scott", "Torres", "Nguyen", "Hill", "Flores",
	"Green", "Adams", "Nelson", "Baker", "Hall", "Rivera", "Campbell", "Mitchell", "Carter", "Roberts",
	"Gomez", "Phillips", "Evans", "Turner", "Diaz", "Parker", "Cruz", "Edwards", "Collins", "Reyes",
	"Stewart", "Morris", "Morales", "Murphy", "Cook", "Rogers", "Gutierrez", "Ortiz", "Morgan", "Cooper",
	"Peterson", "Bailey", "Reed", "Kelly", "Howard", "Ramos", "Kim", "Cox", "Ward", "Richardson",
	"Watson", "Brooks", "Chavez", "Wood", "James", "Bennett", "Gray", "Mendoza", "Ruiz", "Hughes",
	"Price", "Alvarez", "Castillo", "Sanders", "Patel", "Myers", "Long", "Ross", "Foster", "Jimenez",
	"Chavez", "Davidson", "Murphy", "Ferguson", "Wells", "Carlson", "Miller", "Burns", "Bryant", "Hanson",
	"Davidson", "Williamson", "Harrison", "Jenkins", "Chapman", "Hayes", "Holland", "Morrison", "Griffiths",
	"Foster", "Greene", "Curtis", "Chavez", "Robertson", "Franklin", "Gibson", "Chavez", "Stone",
	"Bryan", "Ellis", "Edwards", "Mason", "Hayes", "Bryant", "Jameson", "Hines", "Gibson", "Alexander",
	"Simpson", "Montgomery", "Riley", "Griffith", "Reed", "Douglas", "Hunter", "Barnes", "Kennedy",
	"Douglas", "Hudson", "Carson", "Bishop", "Christensen", "Collier", "Reid", "Stewart", "Curtis",
	"Jordan", "Ross", "Chapman", "McDonald", "Martin", "Simmons", "Vasquez", "Bailey", "Reyes",
	"Gallagher", "Carlson", "Schmidt", "Brooks", "Woods", "Cameron", "Watson", "Walker", "Morrison",
	"Mitchell", "Murray", "Perry", "Ryan", "Sullivan", "Curtis", "Mckinney", "Richards", "Lucas",
	"Craig", "Riley", "James", "Tucker", "Simmons", "Franklin", "Garcia", "Rogers", "Scott", "Payne",
	"Palmer", "Alexander", "Kelly", "Lopez", "Hamilton", "Parker", "Bell", "Scott", "Grant", "Anderson",
	"Gardner", "Harrison", "Patterson", "Jordan", "O'Connor", "Charles", "Bradley", "Simmons", "Bailey",
	"Williamson", "Becker", "Hayward", "Craig", "Harris", "Shaw", "Miller", "Martinez", "Cameron", "Cook",
	"Rice", "Hudson", "Marshall", "King", "Gray", "Stephens", "Matthews", "Stone", "Ford", "Morris",
	"Douglas", "Watts", "Fowler", "Allen", "Simmons", "Barnett", "Graves", "Fleming", "Dixon", "Alexander",
	"Hayes", "Lambert", "Baker", "Parker", "Elliott", "Mason", "Cameron", "Wilson", "Nichols", "Chavez",
}

var initials = []string{
	"A", "A", "B", "B", "C", "C", "D", "D", "E", "E", "E", "E",
	"F", "F", "G", "G", "H", "H", "I", "I", "J", "J", "J", "J",
	"K", "K", "L", "L", "M", "M", "M", "N", "N", "O", "O", "P", "P",
	"Q", "R", "R", "R", "R", "S", "S", "T", "T", "T", "U", "V", "W",
	"X", "Y", "Z",
}
