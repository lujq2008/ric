/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_ReportPeriod_H_
#define	_ReportPeriod_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum ReportPeriod {
	ReportPeriod_ms10	= 0,
	ReportPeriod_ms20	= 1,
	ReportPeriod_ms32	= 2,
	ReportPeriod_ms40	= 3,
	ReportPeriod_ms60	= 4,
	ReportPeriod_ms70	= 5,
	ReportPeriod_ms128	= 6,
	ReportPeriod_ms160	= 7,
	ReportPeriod_ms256	= 8,
	ReportPeriod_ms320	= 9,
	ReportPeriod_ms512	= 10,
	ReportPeriod_ms640	= 11,
	ReportPeriod_ms1024	= 12,
	ReportPeriod_ms2048	= 13,
	ReportPeriod_ms2560	= 14,
	ReportPeriod_ms5120	= 15,
	ReportPeriod_ms10240	= 16
} e_ReportPeriod;

/* ReportPeriod */
typedef long	 ReportPeriod_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_ReportPeriod_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_ReportPeriod;
extern const asn_INTEGER_specifics_t asn_SPC_ReportPeriod_specs_1;
asn_struct_free_f ReportPeriod_free;
asn_struct_print_f ReportPeriod_print;
asn_constr_check_f ReportPeriod_constraint;
ber_type_decoder_f ReportPeriod_decode_ber;
der_type_encoder_f ReportPeriod_encode_der;
xer_type_decoder_f ReportPeriod_decode_xer;
xer_type_encoder_f ReportPeriod_encode_xer;
oer_type_decoder_f ReportPeriod_decode_oer;
oer_type_encoder_f ReportPeriod_encode_oer;
per_type_decoder_f ReportPeriod_decode_uper;
per_type_encoder_f ReportPeriod_encode_uper;
per_type_decoder_f ReportPeriod_decode_aper;
per_type_encoder_f ReportPeriod_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _ReportPeriod_H_ */
#include "asn_internal.h"
