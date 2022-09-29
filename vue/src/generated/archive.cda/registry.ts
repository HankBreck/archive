import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgApproveCda } from "./types/cda/tx";
import { MsgCreateCDA } from "./types/cda/tx";
import { MsgFinalizeCda } from "./types/cda/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/archive.cda.MsgApproveCda", MsgApproveCda],
    ["/archive.cda.MsgCreateCDA", MsgCreateCDA],
    ["/archive.cda.MsgFinalizeCda", MsgFinalizeCda],
    
];

export { msgTypes }