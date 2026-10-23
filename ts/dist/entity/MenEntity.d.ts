import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Men, MenLoadMatch } from '../InkspanDocumentTypes';
declare class MenEntity extends InkspanDocumentEntityBase<Men> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: MenEntity): MenEntity;
    load(this: any, reqmatch?: MenLoadMatch, ctrl?: Control): Promise<MenEntity>;
}
export { MenEntity };
