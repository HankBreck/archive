import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRejectIdentity } from "./types/archive/identity/tx";
import { MsgUpdateOperators } from "./types/archive/identity/tx";
import { MsgRenounceIdentity } from "./types/archive/identity/tx";
import { MsgIssueCertificate } from "./types/archive/identity/tx";
import { MsgRevokeIdentity } from "./types/archive/identity/tx";
import { MsgRegisterIssuer } from "./types/archive/identity/tx";
import { MsgAddIdentityMember } from "./types/archive/identity/tx";
import { MsgAcceptIdentity } from "./types/archive/identity/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/archive.identity.MsgRejectIdentity", MsgRejectIdentity],
    ["/archive.identity.MsgUpdateOperators", MsgUpdateOperators],
    ["/archive.identity.MsgRenounceIdentity", MsgRenounceIdentity],
    ["/archive.identity.MsgIssueCertificate", MsgIssueCertificate],
    ["/archive.identity.MsgRevokeIdentity", MsgRevokeIdentity],
    ["/archive.identity.MsgRegisterIssuer", MsgRegisterIssuer],
    ["/archive.identity.MsgAddIdentityMember", MsgAddIdentityMember],
    ["/archive.identity.MsgAcceptIdentity", MsgAcceptIdentity],
    
];

export { msgTypes }