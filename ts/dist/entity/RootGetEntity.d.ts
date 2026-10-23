import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { RootGet, RootGetLoadMatch } from '../InkspanDocumentTypes';
declare class RootGetEntity extends InkspanDocumentEntityBase<RootGet> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: RootGetEntity): RootGetEntity;
    load(this: any, reqmatch?: RootGetLoadMatch, ctrl?: Control): Promise<RootGetEntity>;
}
export { RootGetEntity };
