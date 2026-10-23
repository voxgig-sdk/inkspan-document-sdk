import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Markdown, MarkdownCreateData } from '../InkspanDocumentTypes';
declare class MarkdownEntity extends InkspanDocumentEntityBase<Markdown> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: MarkdownEntity): MarkdownEntity;
    create(this: any, reqdata?: MarkdownCreateData, ctrl?: Control): Promise<MarkdownEntity>;
}
export { MarkdownEntity };
