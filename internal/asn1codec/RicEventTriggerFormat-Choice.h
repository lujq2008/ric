/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_RicEventTriggerFormat_Choice_H_
#define	_RicEventTriggerFormat_Choice_H_


#include "asn_application.h"

/* Including external dependencies */
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum RicEventTriggerFormat_Choice_PR {
	RicEventTriggerFormat_Choice_PR_NOTHING,	/* No components present */
	RicEventTriggerFormat_Choice_PR_format1,
	RicEventTriggerFormat_Choice_PR_format2,
	RicEventTriggerFormat_Choice_PR_format3,
	RicEventTriggerFormat_Choice_PR_format4,
	RicEventTriggerFormat_Choice_PR_format5,
	RicEventTriggerFormat_Choice_PR_format6,
	RicEventTriggerFormat_Choice_PR_format7,
	RicEventTriggerFormat_Choice_PR_format8,
	RicEventTriggerFormat_Choice_PR_format9,
	RicEventTriggerFormat_Choice_PR_format10
	/* Extensions may appear below */
	
} RicEventTriggerFormat_Choice_PR;

/* Forward declarations */
struct E2SM_TS_EventTriggerDefinitionFormat1;
struct E2SM_TS_EventTriggerDefinitionFormat2;
struct E2SM_TS_EventTriggerDefinitionFormat3;
struct E2SM_TS_EventTriggerDefinitionFormat4;
struct E2SM_TS_EventTriggerDefinitionFormat5;
struct E2SM_TS_EventTriggerDefinitionFormat6;
struct E2SM_TS_EventTriggerDefinitionFormat7;
struct E2SM_TS_EventTriggerDefinitionFormat8;
struct E2SM_TS_EventTriggerDefinitionFormat9;
struct E2SM_TS_EventTriggerDefinitionFormat10;

/* RicEventTriggerFormat-Choice */
typedef struct RicEventTriggerFormat_Choice {
	RicEventTriggerFormat_Choice_PR present;
	union RicEventTriggerFormat_Choice_u {
		struct E2SM_TS_EventTriggerDefinitionFormat1	*format1;
		struct E2SM_TS_EventTriggerDefinitionFormat2	*format2;
		struct E2SM_TS_EventTriggerDefinitionFormat3	*format3;
		struct E2SM_TS_EventTriggerDefinitionFormat4	*format4;
		struct E2SM_TS_EventTriggerDefinitionFormat5	*format5;
		struct E2SM_TS_EventTriggerDefinitionFormat6	*format6;
		struct E2SM_TS_EventTriggerDefinitionFormat7	*format7;
		struct E2SM_TS_EventTriggerDefinitionFormat8	*format8;
		struct E2SM_TS_EventTriggerDefinitionFormat9	*format9;
		struct E2SM_TS_EventTriggerDefinitionFormat10	*format10;
		/*
		 * This type is extensible,
		 * possible extensions are below.
		 */
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} RicEventTriggerFormat_Choice_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_RicEventTriggerFormat_Choice;
extern asn_CHOICE_specifics_t asn_SPC_RicEventTriggerFormat_Choice_specs_1;
extern asn_TYPE_member_t asn_MBR_RicEventTriggerFormat_Choice_1[10];
extern asn_per_constraints_t asn_PER_type_RicEventTriggerFormat_Choice_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _RicEventTriggerFormat_Choice_H_ */
#include "asn_internal.h"
