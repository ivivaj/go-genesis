package consts

// текущая версия
const VERSION = "2.4.1b3"

// чтобы не выдавать одно и тоже голосование
const ASSIGN_TIME = 86400

const DAY = 3600 * 24
const DAY2 = 3600 * 24 * 2

// используется в confirmations
const COUNT_CONFIRMED_NODES = 5
const WAIT_CONFIRMED_NODES = 10

// У скольких нодов должен быть такой же блок как и у нас, чтобы считать, что блок у большей части DC-сети. для get_confirmed_block_id()
const MIN_CONFIRMED_NODES = 0

// примерный текущий крайний блок
const LAST_BLOCK = 330000

// примерный размер блокчейна
const BLOCKCHAIN_SIZE = 77000000

// где лежит блокчейн. для тех, кто не хочет собирать его с нодов
const BLOCKCHAIN_URL = "http://blockchain.daylight.world/blockchain"

// на сколько может бежать время в тр-ии
const MAX_TX_FORW = 0

// тр-ия может блуждать по сети сутки и потом попасть в блок
const MAX_TX_BACK = DAY

const MAX_BLOCK_SIZE = 64<<20

const MAX_TX_SIZE = 32<<20

const MAX_TX_COUNT = 100000

const ERROR_TIME = 1

const DELAY = 5

const ROUND_FIX = 0.00000000001

// таймауты для TCP
const READ_TIMEOUT = 20
const WRITE_TIMEOUT = 20

// дефолтное знаение, со скольким нодами устанавляиваем связь
const OUT_CONNECTIONS = 10



const GAPS_BETWEEN_BLOCKS = 1
const COMMISSION = 1000
const TCP_PORT = "7079"
const MAX_BLOCK_USER_TXS = 100
const RB_BLOCKS_1 = 30
const RB_BLOCKS_2 = 1440
const ALERT_ERROR_TIME = 1

const ALERT_KEY = `30820122300d06092a864886f70d01010105000382010f003082010a0282010100d4a48242d0fb2c7c295bc9c87b1aa0c6d23b5f8cab2ec20c2dfde35513ef6066b92ee3935f9a38100493717b60bb7832411daee02012f44a9f58ac91056b2603661544116bfbc55181e5a693bace5ec9325ba0232b9c9c0a29096569d217243e5bf891cc7fc4bcd2e7d6518acc6f982aaa43a9ed737e3ea2845d6432a823ee5b40d1548f802d0c108bf6e5cb5a4daa7edb48764dcbfa6b7a961208833996cfee265ca2ce2655d444cf3c177b3841b1cc4f3102f89cb2bdb1e5a68eac270506147dd8391b7b3af40a50be13c3970077faffaf98ccc5b8c011146be9c2eb9dfd3454f67a68daaf385d334366d132308bffede27656a515ff69a260bbe2452bd2c30203010001`
const UPD_AND_VER_URL = "http://daylight.world"
const GOOGLE_API_KEY = "AIzaSyBLZlUPgd9uhX05OrsFU68yJOZFrYhZe84"

var LangMap = map[string]int{"en": 1, "ru": 42}

var Countries = []string{"Afghanistan", "Albania", "Algeria", "American Samoa", "Andorra", "Angola", "Anguilla", "Antarctica", "Antigua and Barbuda", "Argentina", "Armenia", "Aruba", "Australia", "Austria", "Azerbaijan", "Bahamas", "Bahrain", "Bangladesh", "Barbados", "Belarus", "Belgium", "Belize", "Benin", "Bermuda", "Bhutan", "Bolivia", "Bosnia and Herzegovina", "Botswana", "Bouvet Island", "Brazil", "British Indian Ocean Territory", "British Virgin Islands", "Brunei", "Bulgaria", "Burkina Faso", "Burundi", "Cambodia", "Cameroon", "Canada", "Cape Verde", "Cayman Islands", "Central African Republic", "Chad", "Chile", "China", "Christmas Island", "Cocos [Keeling] Islands", "Colombia", "Comoros", "Congo [DRC]", "Congo [Republic]", "Cook Islands", "Costa Rica", "Croatia", "Cuba", "Cyprus", "Czech Republic", "Côte d\"Ivoire", "Denmark", "Djibouti", "Dominica", "Dominican Republic", "Ecuador", "Egypt", "El Salvador", "Equatorial Guinea", "Eritrea", "Estonia", "Ethiopia", "Falkland Islands [Islas Malvinas]", "Faroe Islands", "Fiji", "Finland", "France", "French Guiana", "French Polynesia", "French Southern Territories", "Gabon", "Gambia", "Gaza Strip", "Georgia", "Germany", "Ghana", "Gibraltar", "Greece", "Greenland", "Grenada", "Guadeloupe", "Guam", "Guatemala", "Guernsey", "Guinea", "Guinea-Bissau", "Guyana", "Haiti", "Heard Island and McDonald Islands", "Honduras", "Hong Kong", "Hungary", "Iceland", "India", "Indonesia", "Iran", "Iraq", "Ireland", "Isle of Man", "Israel", "Italy", "Jamaica", "Japan", "Jersey", "Jordan", "Kazakhstan", "Kenya", "Kiribati", "Kosovo", "Kuwait", "Kyrgyzstan", "Laos", "Latvia", "Lebanon", "Lesotho", "Liberia", "Libya", "Liechtenstein", "Lithuania", "Luxembourg", "Macau", "Macedonia [FYROM]", "Madagascar", "Malawi", "Malaysia", "Maldives", "Mali", "Malta", "Marshall Islands", "Martinique", "Mauritania", "Mauritius", "Mayotte", "Mexico", "Micronesia", "Moldova", "Monaco", "Mongolia", "Montenegro", "Montserrat", "Morocco", "Mozambique", "Myanmar [Burma]", "Namibia", "Nauru", "Nepal", "Netherlands", "Netherlands Antilles", "New Caledonia", "New Zealand", "Nicaragua", "Niger", "Nigeria", "Niue", "Norfolk Island", "North Korea", "Northern Mariana Islands", "Norway", "Oman", "Pakistan", "Palau", "Palestinian Territories", "Panama", "Papua New Guinea", "Paraguay", "Peru", "Philippines", "Pitcairn Islands", "Poland", "Portugal", "Puerto Rico", "Qatar", "Romania", "Russia", "Rwanda", "Réunion", "Saint Helena", "Saint Kitts and Nevis", "Saint Lucia", "Saint Pierre and Miquelon", "Saint Vincent and the Grenadines", "Samoa", "San Marino", "Saudi Arabia", "Senegal", "Serbia", "Seychelles", "Sierra Leone", "Singapore", "Slovakia", "Slovenia", "Solomon Islands", "Somalia", "South Africa", "South Georgia and the South Sandwich Islands", "South Korea", "Spain", "Sri Lanka", "Sudan", "Suriname", "Svalbard and Jan Mayen", "Swaziland", "Sweden", "Switzerland", "Syria", "São Tomé and Príncipe", "Taiwan", "Tajikistan", "Tanzania", "Thailand", "Timor-Leste", "Togo", "Tokelau", "Tonga", "Trinidad and Tobago", "Tunisia", "Turkey", "Turkmenistan", "Turks and Caicos Islands", "Tuvalu", "U.S. Minor Outlying Islands", "U.S. Virgin Islands", "Uganda", "Ukraine", "United Arab Emirates", "United Kingdom", "United States", "Uruguay", "Uzbekistan", "Vanuatu", "Vatican City", "Venezuela", "Vietnam", "Wallis and Futuna", "Western Sahara", "Yemen", "Zambia", "Zimbabwe"}

var TxTypes = map[int]string{
	1: "FirstBlock",
	2: "DLTTransfer",
	3: "DLTChangeHostVote",
}

func init() {
}