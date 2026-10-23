import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Merge, MergeCreateData } from '../InkspanDocumentTypes';
declare class MergeEntity extends InkspanDocumentEntityBase<Merge> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: MergeEntity): MergeEntity;
    create(this: any, reqdata?: MergeCreateData, ctrl?: Control): Promise<MergeEntity>;
}
export { MergeEntity };
