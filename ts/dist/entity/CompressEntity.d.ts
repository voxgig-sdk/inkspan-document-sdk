import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Compress, CompressCreateData } from '../InkspanDocumentTypes';
declare class CompressEntity extends InkspanDocumentEntityBase<Compress> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: CompressEntity): CompressEntity;
    create(this: any, reqdata?: CompressCreateData, ctrl?: Control): Promise<CompressEntity>;
}
export { CompressEntity };
