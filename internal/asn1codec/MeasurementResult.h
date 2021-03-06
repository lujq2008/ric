/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_MeasurementResult_H_
#define	_MeasurementResult_H_


#include "asn_application.h"

/* Including external dependencies */
#include "MeasurementResultNR.h"
#include "MeasurementResultNeighCells-Choice.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* MeasurementResult */
typedef struct MeasurementResult {
	MeasurementResultNR_t	 measurementResultServingCell;
	MeasurementResultNeighCells_Choice_t	 measurementResultNeighCells;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} MeasurementResult_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_MeasurementResult;
extern asn_SEQUENCE_specifics_t asn_SPC_MeasurementResult_specs_1;
extern asn_TYPE_member_t asn_MBR_MeasurementResult_1[2];

#ifdef __cplusplus
}
#endif

#endif	/* _MeasurementResult_H_ */
#include "asn_internal.h"
