package api

type Response struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

type ChannelListResponse struct {
	Response
	Channels []Channel `json:"channels"`
}
type UserListResponse struct {
	Response
	Users []User `json:"users"`
}

type Channel struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	IsChannel  bool     `json:"is_channel"`
	Created    int64    `json:"created"`
	Creator    string   `json:"creator"`
	IsArchived bool     `json:"is_archived"`
	IsGeneral  bool     `json:"is_general"`
	Members    []string `json:"members"`
	Topic      struct {
		Value   string `json:"value"`
		Creator string `json:"creator"`
		LastSet int64  `json:"last_set"`
	} `json:"topic,omitempty"`
	Purpose struct {
		Value   string `json:"value"`
		Creator string `json:"creator"`
		LastSet int64  `json:"last_set"`
	} `json:"purpose,omitempty"`
	Latest struct {
		Text string  `json:"text"`
		TS   float64 `json:"ts"`
		Type string  `json:"type"`
		User string  `json:"user"`
	} `json:"latest,omitempty"`
	IsMember           bool    `json:"is_member"`
	LastRead           float64 `json:"last_read"`
	UnreadCount        int64   `json:"unread_count"`
	UnreadCountDisplay int64   `json:"unread_count_display"`
}

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Deleted bool   `json:"deleted"`
	Color   string `json:"color"`
	Profile struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		RealName  string `json:"real_name"`
		Email     string `json:"email"`
		Skype     string `json:"skype"`
		Phone     string `json:"phone"`
		*Icons
		/*
			Image24  string `json:"image_24"`
			Image32  string `json:"image_32"`
			Image48  string `json:"image_48"`
			Image72  string `json:"image_72"`
			Image192 string `json:"image_192"`
		*/
	} `json:"profile"`
	IsAdmin           bool `json:"is_admin"`
	IsOwner           bool `json:"is_owner"`
	IsPrimaryOwner    bool `json:"is_primary_owner"`
	IsRestricted      bool `json:"is_restricted"`
	IsUltraRestricted bool `json:"is_ultra_restricted"`
	HasFiles          bool `json:"has_files"`
}

type IM struct {
	ID                 string   `json:"id"`
	IsIM               bool     `json:"is_im"`
	User               string   `json:"user"`
	Created            int64    `json:"created"` // a unix timestamp
	IsUserDeleted      bool     `json:"is_user_deleted"`
	LastRead           float64  `json:"last_read,omitempty"`
	Latest             *Message `json:"latest,omitempty"`
	UnreadCount        int64    `json:"unread_count"`
	UnreadCountDisplay int64    `json:"unread_count_display"`
}

type Group struct {
	Created    int64    `json:"created"`
	Creator    string   `json:"creator"`
	ID         string   `json:"id"`
	IsArchived bool     `json:"is_archived"`
	IsGroup    bool     `json:"is_group"`
	IsOpen     bool     `json:"is_open"`
	LastRead   float64  `json:"last_read"`
	Latest     *Message `json:"latest,omitempty"`
	Members    []string `json:"members"`
	Name       string   `json:"name"`
	Purpose    struct {
		Creator string `json:"creator"`
		LastSet int64  `json:"last_set"`
		Value   string `json:"value"`
	} `json:"purpose"`
	Topic struct {
		Creator string `json:"creator"`
		LastSet int64  `json:"last_set"`
		Value   string `json:"value"`
	} `json:"topic"`
	UnreadCount        int64 `json:"unread_count"`
	UnreadCountDisplay int64 `json:"unread_count_display"`
}

type Team struct {
	ID                string    `json:"id"`
	Domain            string    `json:"domain"`
	EmailDomain       string    `json:"email_domain"`
	Name              string    `json:"name"`
	OverStorageLimit  bool      `json:"over_storage_limit"`
	MsgEditWindowMins int64     `json:"msg_edit_window_mins"`
	Prefs             TeamPrefs `json:"prefs"`
}

type TeamPrefs struct {
	AllowMessageDeletion       bool     `json:"allow_message_deletion"`
	CommandsOnlyRegular        bool     `json:"commands_only_regular"`
	ComplianceExportStart      int64    `json:"compliance_export_start"`
	DefaultChannels            []string `json:"default_channels"`
	DisableBuiltinLoading      bool     `json:"disable_builtin_loading"`
	DisplayRealNames           bool     `json:"display_real_names"`
	DMRetentionDuration        int64    `json:"dm_retention_duration"`
	DMRetentionType            int64    `json:"dm_retention_type"`
	EmojiOnlyAdmins            bool     `json:"emoji_only_admins"`
	GatewayAllowIrcPlain       bool     `json:"gateway_allow_irc_plain"`
	GatewayAllowIrcSSL         bool     `json:"gateway_allow_irc_ssl"`
	GatewayAllowXmppSSL        bool     `json:"gateway_allow_xmpp_ssl"`
	GroupRetentionDuration     int64    `json:"group_retention_duration"`
	GroupRetentionType         int64    `json:"group_retention_type"`
	HideReferers               bool     `json:"hide_referers"`
	LoadingOnlyAdmins          bool     `json:"loading_only_admins"`
	MsgEditWindowMins          int64    `json:"msg_edit_window_mins"`
	RequireAtForMention        int64    `json:"require_at_for_mention"`
	RetentionDuration          int64    `json:"retention_duration"`
	RetentionType              int64    `json:"retention_type"`
	ServicesOnlyAdmin          bool     `json:"services_only_admin"`
	SlackbotResponsesDisabled  bool     `json:"slackbot_responses_disabled"`
	SlackbotResponseOnlyAdmins bool     `json:"slackbot_responses_only_admins"`
	StatsOnlyAdmins            bool     `json:"stats_only_admins"`
	WarnBeforeAtChannel        bool     `json:"warn_before_at_channel"`
	WhoCanArchiveChannels      string   `json:"who_can_archive_channels"`
	WhoCanAtEveryone           string   `json:"who_can_at_everyone"`
	WhoCanCreateChannels       string   `json:"who_can_create_channels"`
	WhoCanCreateGroups         string   `json:"who_can_create_groups"`
	WhoCanKickChannels         string   `json:"who_can_kick_channels"`
	WhoCanKickGroups           string   `json:"who_can_kick_groups"`
	WhoCanPostGeneral          string   `json:"who_can_post_general"`
}

type Bot struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Deleted bool    `json:"deleted"`
	Icons   []Icons `json:"icons"`
}

type Icons struct {
	Image24  string `json:"image_24,omitempty"`
	Image32  string `json:"image_32,omitempty"`
	Image48  string `json:"image_48,omitempty"`
	Image72  string `json:"image_72,omitempty"`
	Image192 string `json:"image_192,omitempty"`
}

type Message struct {
	Text string  `json:text"`
	TS   float64 `json:ts"`
	Type string  `json:type"`
	User string  `json:user"`
}

type File struct {
	// https://api.slack.com/types/file
}
