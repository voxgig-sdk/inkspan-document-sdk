import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Html, HtmlCreateData } from '../InkspanDocumentTypes';
declare class HtmlEntity extends InkspanDocumentEntityBase<Html> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: HtmlEntity): HtmlEntity;
    create(this: any, reqdata?: HtmlCreateData, ctrl?: Control): Promise<HtmlEntity>;
}
export { HtmlEntity };
