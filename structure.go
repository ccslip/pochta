package main

type RequestNormFIO struct {
	OriginalFIO string `json:"original-fio"`
}

type RequestNormAddress struct {
	ID              string `json:"id"`
	OriginalAddress string `json:"original-address"`
}

type RequestNormPhone struct {
	OriginalPhone string `json:"original-phone"`
}

type ResponseNormFIO struct {
	ID          string `json:"id"`           // Уникальный идентификатор
	MiddleName  string `json:"middle-name"`  // Отчество
	Name        string `json:"name"`         // Имя
	OriginalFIO string `json:"original-fio"` // Исходное ФИО (как ввёл пользователь)
	QualityCode string `json:"quality-code"` // Код качества (например, CONFIRMED_MANUALLY)
	Surname     string `json:"surname"`      // Фамилия
	Valid       bool   `json:"valid"`        // Признак валидности
}

type ResponseNormPhone struct {
	ID               string `json:"id"`
	OriginalPhone    string `json:"original-phone"`
	PhoneCityCode    string `json:"phone-city-code"`
	PhoneCountryCode string `json:"phone-country-code"`
	PhoneExtension   string `json:"phone-extension"`
	PhoneNumber      string `json:"phone-number"`
	QualityCode      string `json:"quality-code"`
}

type ResponseNormAddress struct {
	AddressType     string `json:"address-type"`     // Тип адреса (DEFAULT и т.д.)
	Area            string `json:"area"`             // Район/муниципальное образование
	Building        string `json:"building"`         // Строение
	Corpus          string `json:"corpus"`           // Корпус
	Hotel           string `json:"hotel"`            // Гостиница (редко используется)
	House           string `json:"house"`            // Номер дома
	ID              string `json:"id"`               // Уникальный идентификатор
	Index           string `json:"index"`            // Почтовый индекс
	Letter          string `json:"letter"`           // Буква (например, "А" у дома)
	Location        string `json:"location"`         // Населённый пункт
	NumAddressType  string `json:"num-address-type"` // Числовой тип адреса
	OriginalAddress string `json:"original-address"` // Исходный адрес (как ввёл пользователь)
	Place           string `json:"place"`            // Площадь/проспект и т.д.
	QualityCode     string `json:"quality-code"`     // Код качества (GOOD и т.д.)
	Region          string `json:"region"`           // Регион
	Room            string `json:"room"`             // Квартира/помещение
	Slash           string `json:"slash"`            // Дробь (например, "12/3")
	Street          string `json:"street"`           // Улица
	ValidationCode  string `json:"validation-code"`  // Код валидации (CONFIRMED_MANUALLY и т.д.)
}

// Order — основной объект (один элемент массива)
type Order struct {
	AddToMMO             bool   `json:"add-to-mmo,omitempty"`
	AddressTypeTo        string `json:"address-type-to,omitempty"`
	AreaTo               string `json:"area-to,omitempty"`
	BranchName           string `json:"branch-name,omitempty"`
	BuildingTo           string `json:"building-to,omitempty"`
	Comment              string `json:"comment,omitempty"`
	CompletenessChecking bool   `json:"completeness-checking,omitempty"`
	CompulsoryPayment    int    `json:"compulsory-payment,omitempty"`
	CorpusTo             string `json:"corpus-to,omitempty"`
	Courier              bool   `json:"courier,omitempty"`
	DeliveryToDoor       bool   `json:"delivery-to-door,omitempty"`
	DeliveryWithCOD      bool   `json:"delivery-with-cod,omitempty"`
	DimensionType        string `json:"dimension-type,omitempty"`
	EasyReturn           bool   `json:"easy-return,omitempty"`
	Farma                bool   `json:"farma,omitempty"`
	EnvelopeType         string `json:"envelope-type,omitempty"`
	Fragile              bool   `json:"fragile,omitempty"`
	GivenName            string `json:"given-name,omitempty"`
	GroupName            string `json:"group-name,omitempty"`
	HotelTo              string `json:"hotel-to,omitempty"`
	HouseTo              string `json:"house-to,omitempty"`
	IndexTo              int    `json:"index-to,omitempty"`
	InnerNum             string `json:"inner-num,omitempty"`
	InsrValue            int    `json:"insr-value,omitempty"`
	Inventory            bool   `json:"inventory,omitempty"`
	LetterTo             string `json:"letter-to,omitempty"`
	LocationTo           string `json:"location-to,omitempty"`
	ManualAddressInput   bool   `json:"manual-address-input,omitempty"`
	MailCategory         string `json:"mail-category,omitempty"`
	MailDirect           int    `json:"mail-direct,omitempty"`
	MailType             string `json:"mail-type,omitempty"`
	Mass                 int    `json:"mass,omitempty"`
	MiddleName           string `json:"middle-name,omitempty"`
	NoReturn             bool   `json:"no-return,omitempty"`
	NoticePaymentMethod  string `json:"notice-payment-method,omitempty"`
	NumAddressTypeTo     string `json:"num-address-type-to,omitempty"`
	OfficeTo             string `json:"office-to,omitempty"`
	OrderNum             string `json:"order-num,omitempty"`
	Payment              int    `json:"payment,omitempty"`
	PaymentMethod        string `json:"payment-method,omitempty"`
	PlaceTo              string `json:"place-to,omitempty"`
	PostofficeCode       string `json:"postoffice-code,omitempty"`
	PrePostalPreparation bool   `json:"pre-postal-preparation,omitempty"`
	PrepaidAmount        int    `json:"prepaid-amount,omitempty"`
	RecipientName        string `json:"recipient-name,omitempty"`
	RegionTo             string `json:"region-to,omitempty"`
	RoomTo               string `json:"room-to,omitempty"`
	SenderComment        string `json:"sender-comment,omitempty"`
	SenderName           string `json:"sender-name,omitempty"`
	ShelfLifeDays        int    `json:"shelf-life-days,omitempty"`
	SlashTo              string `json:"slash-to,omitempty"`
	SmsNoticeRecipient   int    `json:"sms-notice-recipient,omitempty"`
	StrIndexTo           string `json:"str-index-to,omitempty"`
	StreetTo             string `json:"street-to,omitempty"`
	Surname              string `json:"surname,omitempty"`
	TelAddress           int64  `json:"tel-address,omitempty"`
	TelAddressFrom       int64  `json:"tel-address-from,omitempty"`
	TimeSlotID           int    `json:"time-slot-id,omitempty"`
	Tender               bool   `json:"tender,omitempty"`
	TransportMode        string `json:"transport-mode,omitempty"`
	TransportType        string `json:"transport-type,omitempty"`
	VladenieTo           string `json:"vladenie-to,omitempty"`
	Vsd                  bool   `json:"vsd,omitempty"`
	WithDocuments        bool   `json:"with-documents,omitempty"`
	WithElectronicNotice bool   `json:"with-electronic-notice,omitempty"`
	WithGoods            bool   `json:"with-goods,omitempty"`
	WithOrderOfNotice    bool   `json:"with-order-of-notice,omitempty"`
	WithPackaging        bool   `json:"with-packaging,omitempty"`
	WithSimpleNotice     bool   `json:"with-simple-notice,omitempty"`
	WoMailRank           bool   `json:"wo-mail-rank,omitempty"`
}
type ResponseScriptSiteTo1C struct {
	Contacts Contacts `json:"contacts"`
	Company  Company  `json:"company"`
	Items    []Items  `json:"items"`
	Delivery Delivery `json:"delivery"`
}
type Contacts struct {
	ContragentType     string `json:"contragent__type"`
	Phone              string `json:"phone"`
	Email              string `json:"email"`
	FullName           string `json:"fullname"`
	ContragentTypeName string `json:"_contragent__type"`
}

type Company struct {
	Company string `json:"company"`
	INN     string `json:"inn"`
	Address string `json:"address"`
}

type Items struct {
	Article string  `json:"article"`
	Price   float32 `json:"price"`
	Count   int     `json:"count"`
	Name    string  `json:"name"`
	Weight  int32   `json:"weight"`
	Size    []int32 `json:"size"`
}

type Delivery struct {
	Address string `json:"address"`
	Price   string `json:"price"`
	Index   string `json:"index"`
	City    string `json:"city"`
	Street  string `json:"street"`
	House   string `json:"house"`
}
type ErrorCode struct {
	Code        string `json:"code"`        // Например, "UNDEFINED"
	Description string `json:"description"` // Описание ошибки
	Details     string `json:"details"`     // Детали
	Position    int    `json:"position"`    // Позиция (индекс поля или элемента)
}

// APIError — одна ошибка в массиве errors
type APIError struct {
	ErrorCodes []ErrorCode `json:"error-codes"` // Массив кодов ошибок
	Position   int         `json:"position"`    // Позиция ошибки в запросе
}

// ErrorResponse — полный ответ с ошибками
type ErrorResponse struct {
	Errors    []APIError `json:"errors"`     // Массив ошибок
	ResultIDs []int      `json:"result-ids"` // Массив ID результатов (даже если с ошибкой)
}
