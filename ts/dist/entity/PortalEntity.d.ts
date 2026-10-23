import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Portal, PortalCreateData } from '../InkspanDocumentTypes';
declare class PortalEntity extends InkspanDocumentEntityBase<Portal> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: PortalEntity): PortalEntity;
    create(this: any, reqdata?: PortalCreateData, ctrl?: Control): Promise<PortalEntity>;
}
export { PortalEntity };
