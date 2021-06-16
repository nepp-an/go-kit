package main

import (
    "fmt"
    _ "fmt"
    "reflect"
)

type Enum struct {
    value    int
    show     string
    decision bool
}

type DestinationList struct {
    NORMAL Enum
    INFRA  Enum
    SNS    Enum
}

type IgnoreContentList struct {
    YES Enum
    NO  Enum
}

type ImageList struct {
    YES Enum
    NO  Enum
}

var ImageEnum = ImageList{
    YES: Enum{
        value:    1,
        show:     "YES",
        decision: true,
    },
    NO: Enum{
        value:    0,
        show:     "NO",
        decision: false,
    },
}

type PipelineList struct {
    YES Enum
    NO  Enum
}

var PipelineEnum = PipelineList{
    YES: Enum{
        value:    1,
        show:     "YES",
        decision: true,
    },
    NO: Enum{
        value:    0,
        show:     "NO",
        decision: false,
    },
}

type WhiteList struct {
    YES Enum
    NO  Enum
}

var WhiteListEnum = WhiteList{
    YES: Enum{
        value:    1,
        show:     "YES",
        decision: true,
    },
    NO: Enum{
        value:    0,
        show:     "NO",
        decision: false,
    },
}

type FrequencyList struct {
    YES Enum
    NO  Enum
}

var FrequencyEnum = FrequencyList{
    YES: Enum{
        value:    1,
        show:     "YES",
        decision: true,
    },
    NO: Enum{
        value:    0,
        show:     "NO",
        decision: false,
    },
}

type CrawlerRoleList struct {
    ANALYSIS    Enum
    TRANSCODING Enum
}

var CrawlerRoleEnum = CrawlerRoleList{
    ANALYSIS: Enum{
        value: 1,
        show:  "ANALYSIS",
    },
    TRANSCODING: Enum{
        value: 2,
        show:  "TRANSCODING",
    },
}

type LanguageRemoteDetectList struct {
    YES Enum
    NO  Enum
}

var LanguageRemoteDetectEnum = LanguageRemoteDetectList{
    YES: Enum{
        value:    1,
        show:     "YES",
        decision: true,
    },
    NO: Enum{
        value:    0,
        show:     "NO",
        decision: false,
    },
}

type RepeatingValveList struct {
    YES Enum
    NO  Enum
}

var RepeatingValveEnum = RepeatingValveList{
    YES: Enum{
        value:    1,
        show:     "YES",
        decision: true,
    },
    NO: Enum{
        value:    0,
        show:     "NO",
        decision: false,
    },
}

var IgnoreContentEnum = IgnoreContentList{
    YES: Enum{
        value:    1,
        show:     "YES",
        decision: true,
    },
    NO: Enum{
        value:    0,
        show:     "NO",
        decision: false,
    },
}

var DestinationEnum = DestinationList{
    NORMAL: Enum{
        value: 1,
        show:  "NORMAL",
    },
    INFRA: Enum{
        value: 2,
        show:  "INFRA",
    },
    SNS: Enum{
        value: 3,
        show:  "SNS",
    },
}

type Source struct {
    Value                int
    Show                 string
    Destination          Enum
    IgnoreContent        Enum
    Image                Enum
    Pipeline             Enum
    WhiteList            Enum
    Frequency            Enum
    CrawlerRole          Enum
    RepeatingValve       Enum
    LanguageRemoteDetect Enum
}

type SourceEnum struct {
    NORMAL                               Source
    SNS                                  Source
    CMS                                  Source
    IBM                                  Source
    LOCAL_NEWS                           Source
    DISCOVER                             Source
    RE_TRANSCODING                       Source
    TEST                                 Source
    SOCIAL_DATA_TWITTER                  Source
    IN_CRAWLER                           Source
    STRESS_TEST                          Source
    NLP                                  Source
    INDIA_SNS                            Source
    US_SOCIAL_NETWORK_ANALYSIS           Source
    NEWS_NETWORK_SHOW                    Source
    INDIA_TOP_30                         Source
    SPAM_PORN_PLAYBOY                    Source
    TWITTER_HOME_CRAWLER                 Source
    TWITTER_INFLUENCERS_CRAWLER          Source
    TWITTER_BUZZSUMO_CRAWLER             Source
    TWITTER_FOF_CRAWLER                  Source
    TWITTER_CROWTANGLE_CRAWLER           Source
    FLIPBOARD                            Source
    ZA_CRAWLER                           Source
    NG_CRAWLER                           Source
    KE_CRAWLER                           Source
    GH_CRAWLER                           Source
    TZ_CRAWLER                           Source
    ID_GL_CRAWLER                        Source
    IN_GL_CRAWLER                        Source
    ZA_GL_CRAWLER                        Source
    NG_GL_CRAWLER                        Source
    KE_GL_CRAWLER                        Source
    GH_GL_CRAWLER                        Source
    TZ_GL_CRAWLER                        Source
    IN_UC_CRAWLER                        Source
    EVERGREEN_CRAWLER                    Source
    HINDI_NEWS                           Source
    EXTERNAL_PUBLISHER                   Source
    RU_CRAWLER                           Source
    UA_CRAWLER                           Source
    BY_CRAWLER                           Source
    PK_CRAWLER                           Source
    BD_CRAWLER                           Source
    GLOBAL_CRAWLER                       Source
    US_EN_CRAWLER                        Source
    MY_CRAWLER                           Source
    AR_CRAWLER                           Source
    IN_BN_CRAWLER                        Source
    IN_UR_CRAWLER                        Source
    IN_TA_CRAWLER                        Source
    IN_KN_CRAWLER                        Source
    IN_MR_CRAWLER                        Source
    IN_TE_CRAWLER                        Source
    QUORA_IN_EN                          Source
    QUORA_US                             Source
    QUORA_AF                             Source
    IN_EN_COPYRIGHT                      Source
    ZA_COPYRIGHT                         Source
    ZM_CRAWLER                           Source
    SS_CRAWLER                           Source
    MW_CRAWLER                           Source
    ZW_CRAWLER                           Source
    UG_CRAWLER                           Source
    RW_CRAWLER                           Source
    MG_CRAWLER                           Source
    KT_CRAWLER                           Source
    CM_CRAWLER                           Source
    NE_CRAWLER                           Source
    BF_CRAWLER                           Source
    ML_CRAWLER                           Source
    SN_CRAWLER                           Source
    TD_CRAWLER                           Source
    GN_CRAWLER                           Source
    BI_CRAWLER                           Source
    BJ_CRAWLER                           Source
    CG_CRAWLER                           Source
    QUORA_NG                             Source
    QUORA_GH                             Source
    QUORA_ZA                             Source
    QUORA_KE                             Source
    QUORA_TZ                             Source
    INTEGRATION_CRAWLER_NORMAL           Source
    SELTOOL                              Source
    INTEGRATION_CRAWLER_NORMAL_REPEATING Source
    INTEGRATION_CRAWLER_INFRA            Source
    INTEGRATION_CRAWLER_INFRA_REPEATING  Source
    UG_SW_CRAWLER                        Source
    TZ_SW_CRAWLER                        Source
    NP_CRAWLER                           Source
    LK_CRAWLER                           Source
    AZ_CRAWLER                           Source
    LY_CRAWLER                           Source
    MA_CRAWLER                           Source
    MZ_CRAWLER                           Source
    TR_CRAWLER                           Source
    EG_CRAWLER                           Source
    QUORA_AF_EN                          Source
    GQ_CRAWLER                           Source
    CV_CRAWLER                           Source
    ST_CRAWLER                           Source
    AO_CRAWLER                           Source
    GW_CRAWLER                           Source
    QUORA_AF_FR                          Source
    REDDIT_CRAWLER                       Source
    FLIPBOARD_CRAWLER                    Source
    BOUNCE_CRAWLER                       Source
    DZ_CRAWLER                           Source
    TN_CRAWLER                           Source
    SD_CRAWLER                           Source
    SO_CRAWLER                           Source
    MR_CRAWLER                           Source
    DJ_CRAWLER                           Source
    KM_CRAWLER                           Source
    ENGLISH_COUNTRY_CRAWLER              Source
    AFRICA_ENGLISH_CRAWLER               Source
    QUORA_EN                             Source

    QUORA_HI Source
    QUORA_BN Source
    QUORA_MR Source
    QUORA_TA Source
    QUORA_ID Source

    HA_CRAWLER              Source
    AF_CRAWLER              Source
    XH_CRAWLER              Source
    AM_CRAWLER              Source
    ZU_CRAWLER              Source
    ET_EN_CRAWLER           Source
    GA_CRAWLER              Source
    DZ_FR_CRAWLER           Source
    MA_FR_CRAWLER           Source
    TN_FR_CRAWLER           Source
    ZA_ZULU_CRAWLER         Source
    CONTENT_FACTORY_CRAWLER Source
    NEWS_WITH_TAG_CRAWLER   Source
    NG_YO_CRAWLER           Source
    NEWS_WHIP_CRAWLER       Source
    AFRICA_BLOG_CRAWLER     Source
    PUBLISHER               Source
    NLP_HIGH_LEVEL          Source
    FR_CRAWLER              Source
    FR_CIRCULATION_CRAWLER  Source
    DE_CRAWLER              Source
    QUORA_FR                Source

    THIRD_PARTY_CRAWLER Source
    NEWS_DOG_CRAWLER    Source
    NEWS_POINT_CRAWLER  Source
    UC_NEWS_CRAWLER     Source
    DAILY_HUNT_CRAWLER  Source
    BABE_CRAWLER        Source
}

var SourceList = SourceEnum{
    NORMAL:                               Source{1, "NORMAL", DestinationEnum.NORMAL, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.NO, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SNS:                                  Source{2, "SNS", DestinationEnum.SNS, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    CMS:                                  Source{3, "cms", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    IBM:                                  Source{4, "IBM", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    LOCAL_NEWS:                           Source{5, "Local News", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    DISCOVER:                             Source{6, "DISCOVER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    RE_TRANSCODING:                       Source{7, "RE_TRANSCODING", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.YES},
    TEST:                                 Source{8, "TEST", DestinationEnum.NORMAL, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.NO, WhiteListEnum.NO, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SOCIAL_DATA_TWITTER:                  Source{9, "social-data-twitter", DestinationEnum.INFRA, IgnoreContentEnum.YES, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_CRAWLER:                           Source{10, "IN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    STRESS_TEST:                          Source{11, "STRESS_TEST", DestinationEnum.NORMAL, IgnoreContentEnum.NO, ImageEnum.NO, PipelineEnum.NO, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NLP:                                  Source{12, "NLP", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    INDIA_SNS:                            Source{13, "INDIA_SNS", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    US_SOCIAL_NETWORK_ANALYSIS:           Source{14, "US_SOCIAL_NETWORK_ANALYSIS", DestinationEnum.NORMAL, IgnoreContentEnum.YES, ImageEnum.NO, PipelineEnum.NO, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.ANALYSIS, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NEWS_NETWORK_SHOW:                    Source{15, "NEWS_NETWORK_SHOW", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    INDIA_TOP_30:                         Source{16, "INDIA_TOP_30", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SPAM_PORN_PLAYBOY:                    Source{17, "spam-porn-playboy", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    TWITTER_HOME_CRAWLER:                 Source{18, "TWITTER_HOME_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TWITTER_INFLUENCERS_CRAWLER:          Source{19, "TWITTER_INFLUENCERS_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TWITTER_BUZZSUMO_CRAWLER:             Source{20, "TWITTER_BUZZSUMO_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TWITTER_FOF_CRAWLER:                  Source{21, "TWITTER_FOF_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TWITTER_CROWTANGLE_CRAWLER:           Source{22, "TWITTER_CROWTANGLE_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    FLIPBOARD:                            Source{23, "FLIPBOARD", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    ZA_CRAWLER:                           Source{24, "ZA-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NG_CRAWLER:                           Source{25, "NG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    KE_CRAWLER:                           Source{26, "KE-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    GH_CRAWLER:                           Source{27, "GH-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TZ_CRAWLER:                           Source{28, "TZ-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ID_GL_CRAWLER:                        Source{29, "ID-GL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_GL_CRAWLER:                        Source{30, "IN-GL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ZA_GL_CRAWLER:                        Source{31, "ZA-GL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NG_GL_CRAWLER:                        Source{32, "NG-GL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    KE_GL_CRAWLER:                        Source{33, "KE-GL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    GH_GL_CRAWLER:                        Source{34, "GH-GL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TZ_GL_CRAWLER:                        Source{35, "TZ-GL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_UC_CRAWLER:                        Source{37, "IN-UC-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    EVERGREEN_CRAWLER:                    Source{38, "EVERGREEN_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    HINDI_NEWS:                           Source{39, "HINDI_NEWS", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    EXTERNAL_PUBLISHER:                   Source{40, "EXTERNAL_PUBLISHER", DestinationEnum.NORMAL, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.NO, WhiteListEnum.NO, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    RU_CRAWLER:                           Source{41, "RU-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    UA_CRAWLER:                           Source{42, "UA-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    BY_CRAWLER:                           Source{43, "BY-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    PK_CRAWLER:                           Source{44, "PK-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    BD_CRAWLER:                           Source{45, "BD-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    GLOBAL_CRAWLER:                       Source{46, "GLOBAL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    US_EN_CRAWLER:                        Source{47, "US-EN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    MY_CRAWLER:                           Source{48, "MY-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    AR_CRAWLER:                           Source{49, "AR-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_BN_CRAWLER:                        Source{50, "IN-BN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_UR_CRAWLER:                        Source{51, "IN-UR-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_TA_CRAWLER:                        Source{52, "IN-TA-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_KN_CRAWLER:                        Source{53, "IN-KN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_MR_CRAWLER:                        Source{54, "IN-MR-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_TE_CRAWLER:                        Source{55, "IN-TE-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_IN_EN:                          Source{56, "QUORA-IN-EN", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_US:                             Source{57, "QUORA-US", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_AF:                             Source{58, "QUORA-AF", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_EN_COPYRIGHT:                      Source{59, "IN-EN-COPYRIGHT", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ZA_COPYRIGHT:                         Source{60, "ZA-COPYRIGHT", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ZM_CRAWLER:                           Source{61, "ZM-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SS_CRAWLER:                           Source{62, "SS-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    MW_CRAWLER:                           Source{63, "MW-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ZW_CRAWLER:                           Source{64, "ZW-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    UG_CRAWLER:                           Source{65, "UG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    RW_CRAWLER:                           Source{66, "RW-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    MG_CRAWLER:                           Source{67, "MG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    KT_CRAWLER:                           Source{68, "KT-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    CM_CRAWLER:                           Source{69, "CM-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NE_CRAWLER:                           Source{70, "NE-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    BF_CRAWLER:                           Source{71, "BF-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ML_CRAWLER:                           Source{72, "ML-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SN_CRAWLER:                           Source{73, "SN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TD_CRAWLER:                           Source{74, "TD-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    GN_CRAWLER:                           Source{75, "GN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    BI_CRAWLER:                           Source{76, "BI-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    BJ_CRAWLER:                           Source{77, "BJ-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    CG_CRAWLER:                           Source{78, "CG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_NG:                             Source{79, "QUORA-NG", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_GH:                             Source{80, "QUORA-GH", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_ZA:                             Source{81, "QUORA-ZA", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_KE:                             Source{82, "QUORA-KE", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_TZ:                             Source{83, "QUORA-TZ", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    INTEGRATION_CRAWLER_NORMAL:           Source{84, "INTEGRATION-CRAWLER-NORMAL", DestinationEnum.NORMAL, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SELTOOL:                              Source{85, "SELTOOL", DestinationEnum.NORMAL, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.NO, WhiteListEnum.NO, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    INTEGRATION_CRAWLER_NORMAL_REPEATING: Source{86, "INTEGRATION-CRAWLER-NORMAL-REPEATING", DestinationEnum.NORMAL, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    INTEGRATION_CRAWLER_INFRA:            Source{87, "INTEGRATION-CRAWLER-INFRA", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    INTEGRATION_CRAWLER_INFRA_REPEATING:  Source{88, "INTEGRATION-CRAWLER-INFRA-REPEATING", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    UG_SW_CRAWLER:                        Source{89, "UG-SW-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TZ_SW_CRAWLER:                        Source{90, "TZ-SW-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NP_CRAWLER:                           Source{91, "NP-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    LK_CRAWLER:                           Source{92, "LK-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    AZ_CRAWLER:                           Source{93, "AZ-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    LY_CRAWLER:                           Source{94, "LY-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    MA_CRAWLER:                           Source{95, "MA-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    MZ_CRAWLER:                           Source{96, "MZ-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TR_CRAWLER:                           Source{97, "TR-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    EG_CRAWLER:                           Source{98, "EG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_AF_EN:                          Source{99, "QUORA-AF-EN", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    GQ_CRAWLER:                           Source{100, "GQ-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    CV_CRAWLER:                           Source{101, "CV-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ST_CRAWLER:                           Source{102, "ST-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    AO_CRAWLER:                           Source{103, "AO-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    GW_CRAWLER:                           Source{104, "GW-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_AF_FR:                          Source{105, "QUORA-AF-FR", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    REDDIT_CRAWLER:                       Source{106, "REDDIT-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    FLIPBOARD_CRAWLER:                    Source{107, "FLIPBOARD-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    BOUNCE_CRAWLER:                       Source{108, "BOUNCE-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    DZ_CRAWLER:                           Source{109, "DZ-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TN_CRAWLER:                           Source{110, "TN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SD_CRAWLER:                           Source{111, "SD-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SO_CRAWLER:                           Source{112, "SO-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    MR_CRAWLER:                           Source{113, "MR-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    DJ_CRAWLER:                           Source{114, "DJ-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    KM_CRAWLER:                           Source{115, "KM-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ENGLISH_COUNTRY_CRAWLER:              Source{116, "ENGLISH-COUNTRY-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    AFRICA_ENGLISH_CRAWLER:               Source{117, "AFRICA-ENGLISH-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_EN:                             Source{118, "QUORA-EN", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},

    QUORA_HI: Source{119, "QUORA-HI", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_BN: Source{120, "QUORA-BN", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_MR: Source{121, "QUORA-MR", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_TA: Source{122, "QUORA-TA", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_ID: Source{123, "QUORA-ID", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},

    HA_CRAWLER:              Source{124, "HA-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.YES},
    AF_CRAWLER:              Source{125, "AF-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.YES},
    XH_CRAWLER:              Source{126, "XH-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.YES},
    AM_CRAWLER:              Source{127, "AM-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.YES},
    ZU_CRAWLER:              Source{128, "ZU-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.YES},
    ET_EN_CRAWLER:           Source{129, "ET-EN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    GA_CRAWLER:              Source{130, "GA_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    DZ_FR_CRAWLER:           Source{131, "DZ_FR_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    MA_FR_CRAWLER:           Source{132, "MA_FR_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TN_FR_CRAWLER:           Source{133, "TN_FR_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ZA_ZULU_CRAWLER:         Source{134, "ZA-ZULU-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.YES},
    CONTENT_FACTORY_CRAWLER: Source{135, "CONTENT-FACTORY-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NEWS_WITH_TAG_CRAWLER:   Source{136, "NEWS-WITH-TAG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    NG_YO_CRAWLER:           Source{137, "NG_YO_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    NEWS_WHIP_CRAWLER:       Source{138, "NEWS-WHIP-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    AFRICA_BLOG_CRAWLER:     Source{139, "AFRICA-BLOG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    PUBLISHER:               Source{140, "PUBLISHER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NLP_HIGH_LEVEL:          Source{141, "NLP_HIGH_LEVEL", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    FR_CRAWLER:              Source{142, "FR-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    FR_CIRCULATION_CRAWLER:  Source{143, "FR-CIRCULATION-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    DE_CRAWLER:              Source{144, "DE-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_FR:                Source{145, "QUORA-FR", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},

    THIRD_PARTY_CRAWLER: Source{1000, "THIRD-PARTY-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NEWS_DOG_CRAWLER:    Source{999, "NEWS-DOG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NEWS_POINT_CRAWLER:  Source{998, "NEWS-POINT-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    UC_NEWS_CRAWLER:     Source{997, "UC-NEWS-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    DAILY_HUNT_CRAWLER:  Source{996, "DAILY-HUNT-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    BABE_CRAWLER:        Source{995, "BABE-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
}

const (
    NORMAL = 1
    SNS = 2
    CMS = 3
    IBM = 4
    LOCAL_NEWS = 5
    DISCOVER = 6
    RE_TRANSCODING = 7
    TEST = 8
    SOCIAL_DATA_TWITTER = 9
    IN_CRAWLER = 10
    STRESS_TEST = 11
    NLP = 12
    INDIA_SNS = 13
    US_SOCIAL_NETWORK_ANALYSIS = 14
    NEWS_NETWORK_SHOW = 15
    INDIA_TOP_30 = 16
    SPAM_PORN_PLAYBOY = 17
    TWITTER_HOME_CRAWLER = 18
    TWITTER_INFLUENCERS_CRAWLER = 19
    TWITTER_BUZZSUMO_CRAWLER = 20
    TWITTER_FOF_CRAWLER = 21
    TWITTER_CROWTANGLE_CRAWLER = 22
    FLIPBOARD = 23
    ZA_CRAWLER = 24
    NG_CRAWLER = 25
    KE_CRAWLER = 26
    GH_CRAWLER = 27
    TZ_CRAWLER = 28
    ID_GL_CRAWLER = 29
    IN_GL_CRAWLER = 30
    ZA_GL_CRAWLER = 31
    NG_GL_CRAWLER = 32
    KE_GL_CRAWLER = 33
    GH_GL_CRAWLER = 34
    TZ_GL_CRAWLER = 35
    IN_UC_CRAWLER = 37
    EVERGREEN_CRAWLER = 38
    HINDI_NEWS = 39
    EXTERNAL_PUBLISHER = 40
    RU_CRAWLER = 41
    UA_CRAWLER = 42
    BY_CRAWLER = 43
    PK_CRAWLER = 44
    BD_CRAWLER = 45
    GLOBAL_CRAWLER = 46
    US_EN_CRAWLER = 47
    MY_CRAWLER = 48
    AR_CRAWLER = 49
    IN_BN_CRAWLER = 50
    IN_UR_CRAWLER = 51
    IN_TA_CRAWLER = 52
    IN_KN_CRAWLER = 53
    IN_MR_CRAWLER = 54
    IN_TE_CRAWLER = 55
    QUORA_IN_EN = 56
    QUORA_US = 57
    QUORA_AF = 58
    IN_EN_COPYRIGHT = 59
    ZA_COPYRIGHT = 60
    ZM_CRAWLER = 61
    SS_CRAWLER = 62
    MW_CRAWLER = 63
    ZW_CRAWLER = 64
    UG_CRAWLER = 65
    RW_CRAWLER = 66
    MG_CRAWLER = 67
    KT_CRAWLER = 68
    CM_CRAWLER = 69
    NE_CRAWLER = 70
    BF_CRAWLER = 71
    ML_CRAWLER = 72
    SN_CRAWLER = 73
    TD_CRAWLER = 74
    GN_CRAWLER = 75
    BI_CRAWLER = 76
    BJ_CRAWLER = 77
    CG_CRAWLER = 78
    QUORA_NG = 79
    QUORA_GH = 80
    QUORA_ZA = 81
    QUORA_KE = 82
    QUORA_TZ = 83
    INTEGRATION_CRAWLER_NORMAL = 84
    SELTOOL = 85
    INTEGRATION_CRAWLER_NORMAL_REPEATING = 86
    INTEGRATION_CRAWLER_INFRA = 87
    INTEGRATION_CRAWLER_INFRA_REPEATING = 88
    UG_SW_CRAWLER = 89
    TZ_SW_CRAWLER = 90
    NP_CRAWLER = 91
    LK_CRAWLER = 92
    AZ_CRAWLER = 93
    LY_CRAWLER = 94
    MA_CRAWLER = 95
    MZ_CRAWLER = 96
    TR_CRAWLER = 97
    EG_CRAWLER = 98
    QUORA_AF_EN = 99
    GQ_CRAWLER = 100
    CV_CRAWLER = 101
    ST_CRAWLER = 102
    AO_CRAWLER = 103
    GW_CRAWLER = 104
    QUORA_AF_FR = 105
    REDDIT_CRAWLER = 106
    FLIPBOARD_CRAWLER = 107
    BOUNCE_CRAWLER = 108
    DZ_CRAWLER = 109
    TN_CRAWLER = 110
    SD_CRAWLER = 111
    SO_CRAWLER = 112
    MR_CRAWLER = 113
    DJ_CRAWLER = 114
    KM_CRAWLER = 115
    ENGLISH_COUNTRY_CRAWLER = 116
    AFRICA_ENGLISH_CRAWLER = 117
    QUORA_EN = 118

    QUORA_HI = 119
    QUORA_BN = 120
    QUORA_MR = 121
    QUORA_TA = 122
    QUORA_ID = 123

    HA_CRAWLER = 124
    AF_CRAWLER = 125
    XH_CRAWLER = 126
    AM_CRAWLER = 127
    ZU_CRAWLER = 128
    ET_EN_CRAWLER = 129
    GA_CRAWLER = 130
    DZ_FR_CRAWLER = 131
    MA_FR_CRAWLER = 132
    TN_FR_CRAWLER = 133
    ZA_ZULU_CRAWLER = 134
    CONTENT_FACTORY_CRAWLER = 135
    NEWS_WITH_TAG_CRAWLER = 136
    NG_YO_CRAWLER = 137
    NEWS_WHIP_CRAWLER = 138
    AFRICA_BLOG_CRAWLER = 139
    PUBLISHER = 140
    NLP_HIGH_LEVEL = 141
    FR_CRAWLER = 142
    FR_CIRCULATION_CRAWLER = 143
    DE_CRAWLER = 144
    QUORA_FR = 145

    THIRD_PARTY_CRAWLER = 1000
    NEWS_DOG_CRAWLER = 999
    NEWS_POINT_CRAWLER = 998
    UC_NEWS_CRAWLER = 997
    DAILY_HUNT_CRAWLER = 996
    BABE_CRAWLER = 995
)

var SourceMap = map[int]Source{
    NORMAL:                               Source{1, "NORMAL", DestinationEnum.NORMAL, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.NO, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SNS:                                  Source{2, "SNS", DestinationEnum.SNS, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    CMS:                                  Source{3, "cms", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    IBM:                                  Source{4, "IBM", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    LOCAL_NEWS:                           Source{5, "Local News", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    DISCOVER:                             Source{6, "DISCOVER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    RE_TRANSCODING:                       Source{7, "RE_TRANSCODING", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.YES},
    TEST:                                 Source{8, "TEST", DestinationEnum.NORMAL, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.NO, WhiteListEnum.NO, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SOCIAL_DATA_TWITTER:                  Source{9, "social-data-twitter", DestinationEnum.INFRA, IgnoreContentEnum.YES, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_CRAWLER:                           Source{10, "IN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    STRESS_TEST:                          Source{11, "STRESS_TEST", DestinationEnum.NORMAL, IgnoreContentEnum.NO, ImageEnum.NO, PipelineEnum.NO, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NLP:                                  Source{12, "NLP", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    INDIA_SNS:                            Source{13, "INDIA_SNS", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    US_SOCIAL_NETWORK_ANALYSIS:           Source{14, "US_SOCIAL_NETWORK_ANALYSIS", DestinationEnum.NORMAL, IgnoreContentEnum.YES, ImageEnum.NO, PipelineEnum.NO, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.ANALYSIS, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NEWS_NETWORK_SHOW:                    Source{15, "NEWS_NETWORK_SHOW", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    INDIA_TOP_30:                         Source{16, "INDIA_TOP_30", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SPAM_PORN_PLAYBOY:                    Source{17, "spam-porn-playboy", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    TWITTER_HOME_CRAWLER:                 Source{18, "TWITTER_HOME_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TWITTER_INFLUENCERS_CRAWLER:          Source{19, "TWITTER_INFLUENCERS_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TWITTER_BUZZSUMO_CRAWLER:             Source{20, "TWITTER_BUZZSUMO_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TWITTER_FOF_CRAWLER:                  Source{21, "TWITTER_FOF_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TWITTER_CROWTANGLE_CRAWLER:           Source{22, "TWITTER_CROWTANGLE_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    FLIPBOARD:                            Source{23, "FLIPBOARD", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    ZA_CRAWLER:                           Source{24, "ZA-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NG_CRAWLER:                           Source{25, "NG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    KE_CRAWLER:                           Source{26, "KE-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    GH_CRAWLER:                           Source{27, "GH-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TZ_CRAWLER:                           Source{28, "TZ-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ID_GL_CRAWLER:                        Source{29, "ID-GL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_GL_CRAWLER:                        Source{30, "IN-GL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ZA_GL_CRAWLER:                        Source{31, "ZA-GL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NG_GL_CRAWLER:                        Source{32, "NG-GL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    KE_GL_CRAWLER:                        Source{33, "KE-GL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    GH_GL_CRAWLER:                        Source{34, "GH-GL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TZ_GL_CRAWLER:                        Source{35, "TZ-GL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_UC_CRAWLER:                        Source{37, "IN-UC-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    EVERGREEN_CRAWLER:                    Source{38, "EVERGREEN_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    HINDI_NEWS:                           Source{39, "HINDI_NEWS", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    EXTERNAL_PUBLISHER:                   Source{40, "EXTERNAL_PUBLISHER", DestinationEnum.NORMAL, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.NO, WhiteListEnum.NO, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    RU_CRAWLER:                           Source{41, "RU-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    UA_CRAWLER:                           Source{42, "UA-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    BY_CRAWLER:                           Source{43, "BY-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    PK_CRAWLER:                           Source{44, "PK-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    BD_CRAWLER:                           Source{45, "BD-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    GLOBAL_CRAWLER:                       Source{46, "GLOBAL-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    US_EN_CRAWLER:                        Source{47, "US-EN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    MY_CRAWLER:                           Source{48, "MY-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    AR_CRAWLER:                           Source{49, "AR-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_BN_CRAWLER:                        Source{50, "IN-BN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_UR_CRAWLER:                        Source{51, "IN-UR-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_TA_CRAWLER:                        Source{52, "IN-TA-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_KN_CRAWLER:                        Source{53, "IN-KN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_MR_CRAWLER:                        Source{54, "IN-MR-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_TE_CRAWLER:                        Source{55, "IN-TE-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_IN_EN:                          Source{56, "QUORA-IN-EN", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_US:                             Source{57, "QUORA-US", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_AF:                             Source{58, "QUORA-AF", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    IN_EN_COPYRIGHT:                      Source{59, "IN-EN-COPYRIGHT", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ZA_COPYRIGHT:                         Source{60, "ZA-COPYRIGHT", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ZM_CRAWLER:                           Source{61, "ZM-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SS_CRAWLER:                           Source{62, "SS-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    MW_CRAWLER:                           Source{63, "MW-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ZW_CRAWLER:                           Source{64, "ZW-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    UG_CRAWLER:                           Source{65, "UG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    RW_CRAWLER:                           Source{66, "RW-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    MG_CRAWLER:                           Source{67, "MG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    KT_CRAWLER:                           Source{68, "KT-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    CM_CRAWLER:                           Source{69, "CM-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NE_CRAWLER:                           Source{70, "NE-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    BF_CRAWLER:                           Source{71, "BF-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ML_CRAWLER:                           Source{72, "ML-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SN_CRAWLER:                           Source{73, "SN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TD_CRAWLER:                           Source{74, "TD-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    GN_CRAWLER:                           Source{75, "GN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    BI_CRAWLER:                           Source{76, "BI-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    BJ_CRAWLER:                           Source{77, "BJ-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    CG_CRAWLER:                           Source{78, "CG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_NG:                             Source{79, "QUORA-NG", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_GH:                             Source{80, "QUORA-GH", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_ZA:                             Source{81, "QUORA-ZA", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_KE:                             Source{82, "QUORA-KE", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_TZ:                             Source{83, "QUORA-TZ", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    INTEGRATION_CRAWLER_NORMAL:           Source{84, "INTEGRATION-CRAWLER-NORMAL", DestinationEnum.NORMAL, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SELTOOL:                              Source{85, "SELTOOL", DestinationEnum.NORMAL, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.NO, WhiteListEnum.NO, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    INTEGRATION_CRAWLER_NORMAL_REPEATING: Source{86, "INTEGRATION-CRAWLER-NORMAL-REPEATING", DestinationEnum.NORMAL, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    INTEGRATION_CRAWLER_INFRA:            Source{87, "INTEGRATION-CRAWLER-INFRA", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    INTEGRATION_CRAWLER_INFRA_REPEATING:  Source{88, "INTEGRATION-CRAWLER-INFRA-REPEATING", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    UG_SW_CRAWLER:                        Source{89, "UG-SW-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TZ_SW_CRAWLER:                        Source{90, "TZ-SW-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NP_CRAWLER:                           Source{91, "NP-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    LK_CRAWLER:                           Source{92, "LK-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    AZ_CRAWLER:                           Source{93, "AZ-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    LY_CRAWLER:                           Source{94, "LY-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    MA_CRAWLER:                           Source{95, "MA-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    MZ_CRAWLER:                           Source{96, "MZ-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TR_CRAWLER:                           Source{97, "TR-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    EG_CRAWLER:                           Source{98, "EG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_AF_EN:                          Source{99, "QUORA-AF-EN", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    GQ_CRAWLER:                           Source{100, "GQ-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    CV_CRAWLER:                           Source{101, "CV-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ST_CRAWLER:                           Source{102, "ST-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    AO_CRAWLER:                           Source{103, "AO-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    GW_CRAWLER:                           Source{104, "GW-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_AF_FR:                          Source{105, "QUORA-AF-FR", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    REDDIT_CRAWLER:                       Source{106, "REDDIT-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    FLIPBOARD_CRAWLER:                    Source{107, "FLIPBOARD-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    BOUNCE_CRAWLER:                       Source{108, "BOUNCE-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    DZ_CRAWLER:                           Source{109, "DZ-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TN_CRAWLER:                           Source{110, "TN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SD_CRAWLER:                           Source{111, "SD-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    SO_CRAWLER:                           Source{112, "SO-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    MR_CRAWLER:                           Source{113, "MR-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    DJ_CRAWLER:                           Source{114, "DJ-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    KM_CRAWLER:                           Source{115, "KM-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ENGLISH_COUNTRY_CRAWLER:              Source{116, "ENGLISH-COUNTRY-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    AFRICA_ENGLISH_CRAWLER:               Source{117, "AFRICA-ENGLISH-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_EN:                             Source{118, "QUORA-EN", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},

    QUORA_HI: Source{119, "QUORA-HI", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_BN: Source{120, "QUORA-BN", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_MR: Source{121, "QUORA-MR", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_TA: Source{122, "QUORA-TA", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_ID: Source{123, "QUORA-ID", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},

    HA_CRAWLER:              Source{124, "HA-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.YES},
    AF_CRAWLER:              Source{125, "AF-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.YES},
    XH_CRAWLER:              Source{126, "XH-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.YES},
    AM_CRAWLER:              Source{127, "AM-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.YES},
    ZU_CRAWLER:              Source{128, "ZU-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.YES},
    ET_EN_CRAWLER:           Source{129, "ET-EN-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    GA_CRAWLER:              Source{130, "GA_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    DZ_FR_CRAWLER:           Source{131, "DZ_FR_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    MA_FR_CRAWLER:           Source{132, "MA_FR_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    TN_FR_CRAWLER:           Source{133, "TN_FR_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    ZA_ZULU_CRAWLER:         Source{134, "ZA-ZULU-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.YES},
    CONTENT_FACTORY_CRAWLER: Source{135, "CONTENT-FACTORY-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NEWS_WITH_TAG_CRAWLER:   Source{136, "NEWS-WITH-TAG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    NG_YO_CRAWLER:           Source{137, "NG_YO_CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    NEWS_WHIP_CRAWLER:       Source{138, "NEWS-WHIP-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    AFRICA_BLOG_CRAWLER:     Source{139, "AFRICA-BLOG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    PUBLISHER:               Source{140, "PUBLISHER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NLP_HIGH_LEVEL:          Source{141, "NLP_HIGH_LEVEL", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    FR_CRAWLER:              Source{142, "FR-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    FR_CIRCULATION_CRAWLER:  Source{143, "FR-CIRCULATION-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    DE_CRAWLER:              Source{144, "DE-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    QUORA_FR:                Source{145, "QUORA-FR", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},

    THIRD_PARTY_CRAWLER: Source{1000, "THIRD-PARTY-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NEWS_DOG_CRAWLER:    Source{999, "NEWS-DOG-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    NEWS_POINT_CRAWLER:  Source{998, "NEWS-POINT-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    UC_NEWS_CRAWLER:     Source{997, "UC-NEWS-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.YES, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.NO, LanguageRemoteDetectEnum.NO},
    DAILY_HUNT_CRAWLER:  Source{996, "DAILY-HUNT-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.NO, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
    BABE_CRAWLER:        Source{995, "BABE-CRAWLER", DestinationEnum.INFRA, IgnoreContentEnum.NO, ImageEnum.YES, PipelineEnum.YES, WhiteListEnum.YES, FrequencyEnum.NO, CrawlerRoleEnum.TRANSCODING, RepeatingValveEnum.YES, LanguageRemoteDetectEnum.NO},
}

func main() {
    v := reflect.ValueOf(SourceList)
    count := v.NumField()
    t := reflect.TypeOf(SourceList)
    for i := 0; i < count; i++ {
        fmt.Println(t.Field(i).Name, v.Field(i).FieldByName("Value").Int(), v.Field(i).FieldByName("Destination").FieldByName("show"))

    }
}
