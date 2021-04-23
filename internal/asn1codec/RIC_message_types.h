

#define RIC_UNDEFINED                       -1          //+ Message type is unset or undefined. Newly allocated messages
                                                        //+ by rmr_alloc_msg() will have the type value set to this constant.



/*+
    System Support Message Types
    These message types are used for systems level communications such
    as health checks, alarms and probes.
*/
// Values for system constants are 100-299

#define RIC_HEALTH_CHECK_REQ                100         //+ When received the application is expected to return a response to
                                                        //+ the sender indicating the current "health" of the application.
#define RIC_HEALTH_CHECK_RESP               101         //+ Health responses are sent with a message of this type.

#define RIC_ALARM                           110         //+ Alarm messages with this type are routed to the alarm collection process.
#define RIC_ALARM_QUERY                     111         //+ Unknown meaning

#define RIC_METRICS                         120         //+ This message type causes the message to be routed to the xAPP responsible
                                                        //+ for collecting and redistributing metrics.

// TODO: classify these types for documentation purposes :)
/*+
	Unclassified Message Types
	The following message types have not been clasified into a specific category.
*/

#define RIC_SCTP_CONNECTION_FAILURE         1080
#define RIC_SCTP_CLEAR_ALL                  1090

#define E2_TERM_INIT                        1100
#define E2_TERM_KEEP_ALIVE_REQ              1101
#define E2_TERM_KEEP_ALIVE_RESP             1102

#define RAN_CONNECTED                       1200
#define RAN_RESTARTED                       1210
#define RAN_RECONFIGURED                    1220


#define RIC_ENB_LOAD_INFORMATION            10020

#define RIC_ERROR_INDICATION                10030

#define RIC_SN_STATUS_TRANSFER              10040

#define RIC_UE_CONTEXT_RELEASE              10050

#define RIC_X2_SETUP_REQ                    10060
#define RIC_X2_SETUP_RESP                   10061
#define RIC_X2_SETUP_FAILURE                10062
#define RIC_X2_RESET                        10070
#define RIC_X2_RESET_RESP                   10071

#define RIC_ENB_CONF_UPDATE                 10080
#define RIC_ENB_CONF_UPDATE_ACK             10081
#define RIC_ENB_CONF_UPDATE_FAILURE         10082

#define RIC_RES_STATUS_REQ                  10090
#define RIC_RES_STATUS_RESP                 10091
#define RIC_RES_STATUS_FAILURE              10092

#define RIC_RESOURCE_STATUS_UPDATE          10100

#define RIC_SGNB_ADDITION_REQ               10270
#define RIC_SGNB_ADDITION_ACK               10271
#define RIC_SGNB_ADDITION_REJECT            10272
#define RIC_SGNB_RECONF_COMPLETE            10280

#define RIC_SGNB_MOD_REQUEST                10290
#define RIC_SGNB_MOD_REQUEST_ACK            10291
#define RIC_SGNB_MOD_REQUEST_REJ            10292

#define RIC_SGNB_MOD_REQUIRED               10300
#define RIC_SGNB_MOD_CONFIRM                10301
#define RIC_SGNB_MOD_REFUSE                 10302

#define RIC_SGNB_RELEASE_REQUEST            10310
#define RIC_SGNB_RELEASE_REQUEST_ACK        10311

#define RIC_SGNB_RELEASE_REQUIRED           10320
#define RIC_SGNB_RELEASE_CONFIRM            10321

#define RIC_RRC_TRANSFER                    10350

#define RIC_ENDC_X2_SETUP_REQ               10360
#define RIC_ENDC_X2_SETUP_RESP              10361
#define RIC_ENDC_X2_SETUP_FAILURE           10362

#define RIC_ENDC_CONF_UPDATE                10370
#define RIC_ENDC_CONF_UPDATE_ACK            10371
#define RIC_ENDC_CONF_UPDATE_FAILURE        10372

#define RIC_SECONDARY_RAT_DATA_USAGE_REPORT 10380

#define RIC_GNB_STATUS_INDICATION           10450

#define RIC_E2_SETUP_REQ                    12001
#define RIC_E2_SETUP_RESP                   12002
#define RIC_E2_SETUP_FAILURE                12003
#define RIC_E2_RESET_REQ                    12004
#define RIC_E2_RESET_RESP                   12005
#define RIC_E2_RAN_ERROR_INDICATION         12006
#define RIC_E2_RIC_ERROR_INDICATION         12007
#define RAN_E2_RESET_REQ                    12008
#define RAN_E2_RESET_RESP                   12009



#define RIC_SUB_REQ                         12010
#define RIC_SUB_RESP                        12011
#define RIC_SUB_FAILURE                     12012

#define RIC_SUB_DEL_REQ                     12020
#define RIC_SUB_DEL_RESP                    12021
#define RIC_SUB_DEL_FAILURE                 12022

#define RIC_SERVICE_UPDATE                  12030
#define RIC_SERVICE_UPDATE_ACK              12031
#define RIC_SERVICE_UPDATE_FAILURE          12032

#define RIC_CONTROL_REQ                     12040
#define RIC_CONTROL_ACK                     12041
#define RIC_CONTROL_FAILURE                 12042

#define RIC_INDICATION                      12050

#define RIC_SERVICE_QUERY                   12060

#define DC_ADM_INT_CONTROL                  20000
#define DC_ADM_INT_CONTROL_ACK              20001
#define DC_ADM_GET_POLICY                   20002
#define DC_ADM_GET_POLICY_ACK               20003

#define A1_POLICY_REQ                       20010
#define A1_POLICY_RESP                      20011
#define A1_POLICY_QUERY                     20012


// --- application specific message numbers 30000 - 39999

#define TS_UE_LIST                          30000        // traffic steering
#define TS_QOE_PRED_REQ                     30001
#define TS_QOE_PREDICTION                   30002
#define TS_ANOMALY_UPDATE					30003
#define TS_ANOMALY_ACK						30004


#define MC_REPORT                           30010        // Measurement campaign xApp reports


// ---- these are old (release 0) definitions and should not be used ------

/*
    E2 Related messages  should be in the range 10000 to 99999
*/

#define RIC_X2_SETUP                        10000        //-
#define RIC_X2_RESPONSE                     10001        //-
#define RIC_X2_RESOURCE_STATUS_REQUEST      10002        //-
#define RIC_X2_RESOURCE_STATUS_RESPONSE     10003        //-
#define RIC_X2_LOAD_INFORMATION             10004        //-
#define RIC_E2_TERMINATION_HC_REQUEST       10005        //-
#define RIC_E2_TERMINATION_HC_RESPONSE      10006        //-
#define RIC_E2_MANAGER_HC_REQUEST           10007        //-
#define RIC_E2_MANAGER_HC_RESPONSE          10008        //-


/*
    A1 Related messages should be in the range 100000 to 999999
*/
#define RIC_CONTROL_XAPP_CONFIG_REQUEST     100000      //-
#define RIC_CONTROL_XAPP_CONFIG_RESPONSE    100001      //-
