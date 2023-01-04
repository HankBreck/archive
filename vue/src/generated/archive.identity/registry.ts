import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgIssueCertificate } from "./types/archive/identity/tx";
import { MsgRenounceIdentity } from "./types/archive/identity/tx";
import { MsgAddIdentityMember } from "./types/archive/identity/tx";
import { MsgUpdateMembers } from "./types/archive/identity/tx";
import { MsgUpdateOperators } from "./types/archive/identity/tx";
import { MsgRejectIdentity } from "./types/archive/identity/tx";
import { MsgRevokeIdentity } from "./types/archive/identity/tx";
import { MsgRegisterIssuer } from "./types/archive/identity/tx";
import { MsgAcceptIdentity } from "./types/archive/identity/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/archive.identity.MsgIssueCertificate", MsgIssueCertificate],
    ["/archive.identity.MsgRenounceIdentity", MsgRenounceIdentity],
    ["/archive.identity.MsgAddIdentityMember", MsgAddIdentityMember],
    ["/archive.identity.MsgUpdateMembers", MsgUpdateMembers],
    ["/archive.identity.MsgUpdateOperators", MsgUpdateOperators],
    ["/archive.identity.MsgRejectIdentity", MsgRejectIdentity],
    ["/archive.identity.MsgRevokeIdentity", MsgRevokeIdentity],
    ["/archive.identity.MsgRegisterIssuer", MsgRegisterIssuer],
    ["/archive.identity.MsgAcceptIdentity", MsgAcceptIdentity],
    
];

export { msgTypes }