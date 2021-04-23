/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_MeasurementResultNR_H_
#define	_MeasurementResultNR_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct MeasurementQuantityResult;

/* MeasurementResultNR */
typedef struct MeasurementResultNR {
	long	 pCI;
	struct MeasurementQuantityResult	*based_SSB;	/* OPTIONAL */
	struct MeasurementQuantityResult	*based_CSI_RS;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} MeasurementResultNR_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_MeasurementResultNR;
extern asn_SEQUENCE_specifics_t asn_SPC_MeasurementResultNR_specs_1;
extern asn_TYPE_member_t asn_MBR_MeasurementResultNR_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _MeasurementResultNR_H_ */
#include "asn_internal.h"
