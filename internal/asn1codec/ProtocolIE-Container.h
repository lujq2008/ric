/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-Containers"
 * 	found in "../nric_e2ap_v01.01.asn1"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_ProtocolIE_Container_H_
#define	_ProtocolIE_Container_H_


#include "asn_application.h"

/* Including external dependencies */
#include "asn_SEQUENCE_OF.h"
#include "constr_SEQUENCE_OF.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct RICsubscriptionRequest_IEs;
struct RICsubscriptionResponse_IEs;
struct RICsubscriptionFailure_IEs;
struct RICsubscriptionDeleteRequest_IEs;
struct RICsubscriptionDeleteResponse_IEs;
struct RICsubscriptionDeleteFailure_IEs;
struct RICindication_IEs;
struct RICcontrolRequest_IEs;
struct RICcontrolAcknowledge_IEs;
struct RICcontrolFailure_IEs;
struct ErrorIndication_IEs;
struct E2setupRequestIEs;
struct E2setupResponseIEs;
struct E2setupFailureIEs;
struct E2connectionUpdate_IEs;
struct E2connectionUpdateAck_IEs;
struct E2connectionUpdateFailure_IEs;
struct E2nodeConfigurationUpdate_IEs;
struct E2nodeConfigurationUpdateAcknowledge_IEs;
struct E2nodeConfigurationUpdateFailure_IEs;
struct ResetRequestIEs;
struct ResetResponseIEs;
struct RICserviceUpdate_IEs;
struct RICserviceUpdateAcknowledge_IEs;
struct RICserviceUpdateFailure_IEs;
struct RICserviceQuery_IEs;

/* ProtocolIE-Container */
typedef struct ProtocolIE_Container_1724P0 {
	A_SEQUENCE_OF(struct RICsubscriptionRequest_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P0_t;
typedef struct ProtocolIE_Container_1724P1 {
	A_SEQUENCE_OF(struct RICsubscriptionResponse_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P1_t;
typedef struct ProtocolIE_Container_1724P2 {
	A_SEQUENCE_OF(struct RICsubscriptionFailure_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P2_t;
typedef struct ProtocolIE_Container_1724P3 {
	A_SEQUENCE_OF(struct RICsubscriptionDeleteRequest_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P3_t;
typedef struct ProtocolIE_Container_1724P4 {
	A_SEQUENCE_OF(struct RICsubscriptionDeleteResponse_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P4_t;
typedef struct ProtocolIE_Container_1724P5 {
	A_SEQUENCE_OF(struct RICsubscriptionDeleteFailure_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P5_t;
typedef struct ProtocolIE_Container_1724P6 {
	A_SEQUENCE_OF(struct RICindication_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P6_t;
typedef struct ProtocolIE_Container_1724P7 {
	A_SEQUENCE_OF(struct RICcontrolRequest_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P7_t;
typedef struct ProtocolIE_Container_1724P8 {
	A_SEQUENCE_OF(struct RICcontrolAcknowledge_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P8_t;
typedef struct ProtocolIE_Container_1724P9 {
	A_SEQUENCE_OF(struct RICcontrolFailure_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P9_t;
typedef struct ProtocolIE_Container_1724P10 {
	A_SEQUENCE_OF(struct ErrorIndication_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P10_t;
typedef struct ProtocolIE_Container_1724P11 {
	A_SEQUENCE_OF(struct E2setupRequestIEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P11_t;
typedef struct ProtocolIE_Container_1724P12 {
	A_SEQUENCE_OF(struct E2setupResponseIEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P12_t;
typedef struct ProtocolIE_Container_1724P13 {
	A_SEQUENCE_OF(struct E2setupFailureIEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P13_t;
typedef struct ProtocolIE_Container_1724P14 {
	A_SEQUENCE_OF(struct E2connectionUpdate_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P14_t;
typedef struct ProtocolIE_Container_1724P15 {
	A_SEQUENCE_OF(struct E2connectionUpdateAck_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P15_t;
typedef struct ProtocolIE_Container_1724P16 {
	A_SEQUENCE_OF(struct E2connectionUpdateFailure_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P16_t;
typedef struct ProtocolIE_Container_1724P17 {
	A_SEQUENCE_OF(struct E2nodeConfigurationUpdate_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P17_t;
typedef struct ProtocolIE_Container_1724P18 {
	A_SEQUENCE_OF(struct E2nodeConfigurationUpdateAcknowledge_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P18_t;
typedef struct ProtocolIE_Container_1724P19 {
	A_SEQUENCE_OF(struct E2nodeConfigurationUpdateFailure_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P19_t;
typedef struct ProtocolIE_Container_1724P20 {
	A_SEQUENCE_OF(struct ResetRequestIEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P20_t;
typedef struct ProtocolIE_Container_1724P21 {
	A_SEQUENCE_OF(struct ResetResponseIEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P21_t;
typedef struct ProtocolIE_Container_1724P22 {
	A_SEQUENCE_OF(struct RICserviceUpdate_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P22_t;
typedef struct ProtocolIE_Container_1724P23 {
	A_SEQUENCE_OF(struct RICserviceUpdateAcknowledge_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P23_t;
typedef struct ProtocolIE_Container_1724P24 {
	A_SEQUENCE_OF(struct RICserviceUpdateFailure_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P24_t;
typedef struct ProtocolIE_Container_1724P25 {
	A_SEQUENCE_OF(struct RICserviceQuery_IEs) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ProtocolIE_Container_1724P25_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P0;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P0_specs_1;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P0_1[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P0_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P1;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P1_specs_3;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P1_3[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P1_constr_3;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P2;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P2_specs_5;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P2_5[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P2_constr_5;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P3;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P3_specs_7;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P3_7[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P3_constr_7;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P4;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P4_specs_9;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P4_9[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P4_constr_9;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P5;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P5_specs_11;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P5_11[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P5_constr_11;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P6;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P6_specs_13;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P6_13[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P6_constr_13;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P7;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P7_specs_15;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P7_15[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P7_constr_15;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P8;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P8_specs_17;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P8_17[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P8_constr_17;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P9;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P9_specs_19;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P9_19[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P9_constr_19;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P10;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P10_specs_21;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P10_21[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P10_constr_21;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P11;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P11_specs_23;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P11_23[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P11_constr_23;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P12;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P12_specs_25;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P12_25[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P12_constr_25;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P13;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P13_specs_27;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P13_27[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P13_constr_27;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P14;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P14_specs_29;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P14_29[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P14_constr_29;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P15;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P15_specs_31;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P15_31[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P15_constr_31;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P16;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P16_specs_33;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P16_33[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P16_constr_33;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P17;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P17_specs_35;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P17_35[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P17_constr_35;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P18;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P18_specs_37;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P18_37[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P18_constr_37;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P19;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P19_specs_39;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P19_39[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P19_constr_39;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P20;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P20_specs_41;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P20_41[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P20_constr_41;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P21;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P21_specs_43;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P21_43[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P21_constr_43;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P22;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P22_specs_45;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P22_45[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P22_constr_45;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P23;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P23_specs_47;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P23_47[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P23_constr_47;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P24;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P24_specs_49;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P24_49[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P24_constr_49;
extern asn_TYPE_descriptor_t asn_DEF_ProtocolIE_Container_1724P25;
extern asn_SET_OF_specifics_t asn_SPC_ProtocolIE_Container_1724P25_specs_51;
extern asn_TYPE_member_t asn_MBR_ProtocolIE_Container_1724P25_51[1];
extern asn_per_constraints_t asn_PER_type_ProtocolIE_Container_1724P25_constr_51;

#ifdef __cplusplus
}
#endif

#endif	/* _ProtocolIE_Container_H_ */
#include "asn_internal.h"