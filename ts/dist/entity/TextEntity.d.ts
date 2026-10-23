import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Text, TextCreateData } from '../InkspanDocumentTypes';
declare class TextEntity extends InkspanDocumentEntityBase<Text> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: TextEntity): TextEntity;
    create(this: any, reqdata?: TextCreateData, ctrl?: Control): Promise<TextEntity>;
}
export { TextEntity };
