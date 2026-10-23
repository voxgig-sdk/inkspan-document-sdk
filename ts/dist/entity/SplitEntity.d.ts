import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Split, SplitCreateData } from '../InkspanDocumentTypes';
declare class SplitEntity extends InkspanDocumentEntityBase<Split> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: SplitEntity): SplitEntity;
    create(this: any, reqdata?: SplitCreateData, ctrl?: Control): Promise<SplitEntity>;
}
export { SplitEntity };
