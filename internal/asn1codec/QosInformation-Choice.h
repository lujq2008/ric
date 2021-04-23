/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_QosInformation_Choice_H_
#define	_QosInformation_Choice_H_


#include "asn_application.h"

/* Including external dependencies */
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum QosInformation_Choice_PR {
	QosInformation_Choice_PR_NOTHING,	/* No components present */
	QosInformation_Choice_PR_e_UTRANQos,
	QosInformation_Choice_PR_dRBInformation
	/* Extensions may appear below */
	
} QosInformation_Choice_PR;

/* Forward declarations */
struct E_UTRANQoS;
struct DRBQoS;

/* QosInformation-Choice */
typedef struct QosInformation_Choice {
	QosInformation_Choice_PR present;
	union QosInformation_Choice_u {
		struct E_UTRANQoS	*e_UTRANQos;
		struct DRBQoS	*dRBInformation;
		/*
		 * This type is extensible,
		 * possible extensions are below.
		 */
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} QosInformation_Choice_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_QosInformation_Choice;
extern asn_CHOICE_specifics_t asn_SPC_QosInformation_Choice_specs_1;
extern asn_TYPE_member_t asn_MBR_QosInformation_Choice_1[2];
extern asn_per_constraints_t asn_PER_type_QosInformation_Choice_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _QosInformation_Choice_H_ */
#include "asn_internal.h"
