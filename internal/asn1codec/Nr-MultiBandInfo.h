/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_Nr_MultiBandInfo_H_
#define	_Nr_MultiBandInfo_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "asn_SEQUENCE_OF.h"
#include "constr_SEQUENCE_OF.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct Nr_Ns_Pmax;

/* Nr-MultiBandInfo */
typedef struct Nr_MultiBandInfo {
	long	*frequencyBandIndicatorNr;	/* OPTIONAL */
	struct Nr_MultiBandInfo__nR_NS_Pmax_List {
		A_SEQUENCE_OF(struct Nr_Ns_Pmax) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} *nR_NS_Pmax_List;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} Nr_MultiBandInfo_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_Nr_MultiBandInfo;
extern asn_SEQUENCE_specifics_t asn_SPC_Nr_MultiBandInfo_specs_1;
extern asn_TYPE_member_t asn_MBR_Nr_MultiBandInfo_1[2];

#ifdef __cplusplus
}
#endif

#endif	/* _Nr_MultiBandInfo_H_ */
#include "asn_internal.h"
