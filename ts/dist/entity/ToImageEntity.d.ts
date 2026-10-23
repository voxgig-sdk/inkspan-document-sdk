import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { ToImage, ToImageCreateData } from '../InkspanDocumentTypes';
declare class ToImageEntity extends InkspanDocumentEntityBase<ToImage> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: ToImageEntity): ToImageEntity;
    create(this: any, reqdata?: ToImageCreateData, ctrl?: Control): Promise<ToImageEntity>;
}
export { ToImageEntity };
