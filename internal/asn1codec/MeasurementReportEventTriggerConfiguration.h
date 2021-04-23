/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_MeasurementReportEventTriggerConfiguration_H_
#define	_MeasurementReportEventTriggerConfiguration_H_


#include "asn_application.h"

/* Including external dependencies */
#include "asn_SEQUENCE_OF.h"
#include "constr_SEQUENCE_OF.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct A1ReportConfiguration;
struct A2ReportConfiguration;
struct A3ReportConfiguration;
struct A4ReportConfiguration;
struct A5ReportConfiguration;
struct A6ReportConfiguration;
struct B1ReportConfiguration;
struct B2ReportConfiguration;

/* MeasurementReportEventTriggerConfiguration */
typedef struct MeasurementReportEventTriggerConfiguration {
	struct MeasurementReportEventTriggerConfiguration__a1ReportConfiguration_List {
		A_SEQUENCE_OF(struct A1ReportConfiguration) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} a1ReportConfiguration_List;
	struct MeasurementReportEventTriggerConfiguration__a2ReportConfiguration_List {
		A_SEQUENCE_OF(struct A2ReportConfiguration) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} a2ReportConfiguration_List;
	struct MeasurementReportEventTriggerConfiguration__a3ReportConfiguration_List {
		A_SEQUENCE_OF(struct A3ReportConfiguration) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} a3ReportConfiguration_List;
	struct MeasurementReportEventTriggerConfiguration__a4ReportConfiguration_List {
		A_SEQUENCE_OF(struct A4ReportConfiguration) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} a4ReportConfiguration_List;
	struct MeasurementReportEventTriggerConfiguration__a5ReportConfiguration_List {
		A_SEQUENCE_OF(struct A5ReportConfiguration) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} a5ReportConfiguration_List;
	struct MeasurementReportEventTriggerConfiguration__a6ReportConfiguration_List {
		A_SEQUENCE_OF(struct A6ReportConfiguration) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} a6ReportConfiguration_List;
	struct MeasurementReportEventTriggerConfiguration__b1ReportConfiguration_List {
		A_SEQUENCE_OF(struct B1ReportConfiguration) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} b1ReportConfiguration_List;
	struct MeasurementReportEventTriggerConfiguration__b2ReportConfiguration_List {
		A_SEQUENCE_OF(struct B2ReportConfiguration) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} b2ReportConfiguration_List;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} MeasurementReportEventTriggerConfiguration_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_MeasurementReportEventTriggerConfiguration;
extern asn_SEQUENCE_specifics_t asn_SPC_MeasurementReportEventTriggerConfiguration_specs_1;
extern asn_TYPE_member_t asn_MBR_MeasurementReportEventTriggerConfiguration_1[8];

#ifdef __cplusplus
}
#endif

#endif	/* _MeasurementReportEventTriggerConfiguration_H_ */
#include "asn_internal.h"
