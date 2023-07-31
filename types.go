package gosdk

type ComponentTextSize string

const (
	ComponentTextSizeS ComponentTextSize = "S"
	ComponentTextSizeM ComponentTextSize = "M"
	ComponentTextSizeL ComponentTextSize = "L"
)

type ComponentTextColor string

const (
	ComponentTextColorNormal  ComponentTextColor = "NORMAL"
	ComponentTextColorMuted   ComponentTextColor = "MUTED"
	ComponentTextColorSuccess ComponentTextColor = "SUCCESS"
	ComponentTextColorWarning ComponentTextColor = "WARNING"
	ComponentTextColorError   ComponentTextColor = "ERROR"
)

type ComponentSpacerSize string

const (
	ComponentSpacerSizeXS ComponentSpacerSize = "XS"
	ComponentSpacerSizeS  ComponentSpacerSize = "S"
	ComponentSpacerSizeM  ComponentSpacerSize = "M"
	ComponentSpacerSizeL  ComponentSpacerSize = "L"
	ComponentSpacerSizeXL ComponentSpacerSize = "XL"
)

type ComponentDividerSpacingSize string

const (
	ComponentDividerSpacingSizeXS ComponentDividerSpacingSize = "XS"
	ComponentDividerSpacingSizeS  ComponentDividerSpacingSize = "S"
	ComponentDividerSpacingSizeM  ComponentDividerSpacingSize = "M"
	ComponentDividerSpacingSizeL  ComponentDividerSpacingSize = "L"
	ComponentDividerSpacingSizeXL ComponentDividerSpacingSize = "XL"
)

type ComponentBadgeColor string

const (
	ComponentBadgeColorGrey   ComponentBadgeColor = "GREY"
	ComponentBadgeColorGreen  ComponentBadgeColor = "GREEN"
	ComponentBadgeColorYellow ComponentBadgeColor = "YELLOW"
	ComponentBadgeColorRed    ComponentBadgeColor = "RED"
	ComponentBadgeColorBlue   ComponentBadgeColor = "BLUE"
)

type ComponentText struct {
	TextSize  *ComponentTextSize  `json:"textSize,omitempty"`
	TextColor *ComponentTextColor `json:"textColor,omitempty"`
	Text      string              `json:"text"`
}

type ComponentDivider struct {
	DividerSpacingSize *ComponentDividerSpacingSize `json:"dividerSpacingSize,omitempty"`
}

type ComponentLinkButton struct {
	LinkButtonURL   string `json:"linkButtonUrl"`
	LinkButtonLabel string `json:"linkButtonLabel"`
}

type ComponentSpacer struct {
	SpacerSize ComponentSpacerSize `json:"spacerSize"`
}

type ComponentBadge struct {
	BadgeLabel string               `json:"badgeLabel"`
	BadgeColor *ComponentBadgeColor `json:"badgeColor,omitempty"`
}

type ComponentCopyButton struct {
	CopyButtonValue        string  `json:"copyButtonValue"`
	CopyButtonTooltipLabel *string `json:"copyButtonTooltipLabel,omitempty"`
}

type ComponentRowContentUnionInput struct {
	ComponentText       *ComponentText       `json:"componentText,omitempty"`
	ComponentDivider    *ComponentDivider    `json:"componentDivider,omitempty"`
	ComponentLinkButton *ComponentLinkButton `json:"componentLinkButton,omitempty"`
	ComponentSpacer     *ComponentSpacer     `json:"componentSpacer,omitempty"`
	ComponentBadge      *ComponentBadge      `json:"componentBadge,omitempty"`
	ComponentCopyButton *ComponentCopyButton `json:"componentCopyButton,omitempty"`
}

type ComponentRow struct {
	RowMainContent  []ComponentRowContentUnionInput `json:"rowMainContent"`
	RowAsideContent []ComponentRowContentUnionInput `json:"rowAsideContent"`
}

type ComponentContainerContentUnionInput struct {
	ComponentText       *ComponentText       `json:"componentText,omitempty"`
	ComponentDivider    *ComponentDivider    `json:"componentDivider,omitempty"`
	ComponentLinkButton *ComponentLinkButton `json:"componentLinkButton,omitempty"`
	ComponentSpacer     *ComponentSpacer     `json:"componentSpacer,omitempty"`
	ComponentBadge      *ComponentBadge      `json:"componentBadge,omitempty"`
	ComponentCopyButton *ComponentCopyButton `json:"componentCopyButton,omitempty"`
	ComponentRow        *ComponentRow        `json:"componentRow,omitempty"`
}

type ComponentContainer struct {
	ContainerContent []ComponentContainerContentUnionInput `json:"containerContent"`
}

type Component struct {
	ComponentText       *ComponentText       `json:"componentText,omitempty"`
	ComponentDivider    *ComponentDivider    `json:"componentDivider,omitempty"`
	ComponentLinkButton *ComponentLinkButton `json:"componentLinkButton,omitempty"`
	ComponentSpacer     *ComponentSpacer     `json:"componentSpacer,omitempty"`
	ComponentBadge      *ComponentBadge      `json:"componentBadge,omitempty"`
	ComponentCopyButton *ComponentCopyButton `json:"componentCopyButton,omitempty"`
	ComponentRow        *ComponentRow        `json:"componentRow,omitempty"`
	ComponentContainer  *ComponentContainer  `json:"componentContainer,omitempty"`
}

type CustomerCard struct {
	Key               string      `json:"key"`
	TimeToLiveSeconds *int        `json:"timeToLiveSeconds,omitempty"`
	Components        []Component `json:"components,omitempty"`
}
