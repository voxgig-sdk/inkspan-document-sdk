import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Inbound, InboundCreateData } from '../InkspanDocumentTypes';
declare class InboundEntity extends InkspanDocumentEntityBase<Inbound> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: InboundEntity): InboundEntity;
    create(this: any, reqdata?: InboundCreateData, ctrl?: Control): Promise<InboundEntity>;
}
export { InboundEntity };
