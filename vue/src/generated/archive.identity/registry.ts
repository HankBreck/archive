import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgIssueCertificate } from "./types/archive/identity/tx";
import { MsgRegisterIssuer } from "./types/archive/identity/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/archive.identity.MsgIssueCertificate", MsgIssueCertificate],
    ["/archive.identity.MsgRegisterIssuer", MsgRegisterIssuer],
    
];

export { msgTypes }