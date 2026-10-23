import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Template, TemplateCreateData } from '../InkspanDocumentTypes';
declare class TemplateEntity extends InkspanDocumentEntityBase<Template> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: TemplateEntity): TemplateEntity;
    create(this: any, reqdata?: TemplateCreateData, ctrl?: Control): Promise<TemplateEntity>;
}
export { TemplateEntity };
