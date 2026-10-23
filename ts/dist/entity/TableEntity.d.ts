import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Table, TableCreateData } from '../InkspanDocumentTypes';
declare class TableEntity extends InkspanDocumentEntityBase<Table> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: TableEntity): TableEntity;
    create(this: any, reqdata?: TableCreateData, ctrl?: Control): Promise<TableEntity>;
}
export { TableEntity };
