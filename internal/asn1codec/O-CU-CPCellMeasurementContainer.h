/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_O_CU_CPCellMeasurementContainer_H_
#define	_O_CU_CPCellMeasurementContainer_H_


#include "asn_application.h"

/* Including external dependencies */
#include "asn_SEQUENCE_OF.h"
#include "constr_SEQUENCE_OF.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct O_Cu_CpCellMeasurement;

/* O-CU-CPCellMeasurementContainer */
typedef struct O_CU_CPCellMeasurementContainer {
	struct O_CU_CPCellMeasurementContainer__o_CU_CPCellMeasurement_List {
		A_SEQUENCE_OF(struct O_Cu_CpCellMeasurement) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} o_CU_CPCellMeasurement_List;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} O_CU_CPCellMeasurementContainer_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_O_CU_CPCellMeasurementContainer;
extern asn_SEQUENCE_specifics_t asn_SPC_O_CU_CPCellMeasurementContainer_specs_1;
extern asn_TYPE_member_t asn_MBR_O_CU_CPCellMeasurementContainer_1[1];

#ifdef __cplusplus
}
#endif

#endif	/* _O_CU_CPCellMeasurementContainer_H_ */
#include "asn_internal.h"
