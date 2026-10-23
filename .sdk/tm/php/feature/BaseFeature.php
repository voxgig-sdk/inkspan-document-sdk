<?php
declare(strict_types=1);

// InkspanDocument SDK base feature

class InkspanDocumentBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(InkspanDocumentContext $ctx, array $options): void {}
    public function PostConstruct(InkspanDocumentContext $ctx): void {}
    public function PostConstructEntity(InkspanDocumentContext $ctx): void {}
    public function SetData(InkspanDocumentContext $ctx): void {}
    public function GetData(InkspanDocumentContext $ctx): void {}
    public function GetMatch(InkspanDocumentContext $ctx): void {}
    public function SetMatch(InkspanDocumentContext $ctx): void {}
    public function PrePoint(InkspanDocumentContext $ctx): void {}
    public function PreSpec(InkspanDocumentContext $ctx): void {}
    public function PreRequest(InkspanDocumentContext $ctx): void {}
    public function PreResponse(InkspanDocumentContext $ctx): void {}
    public function PreResult(InkspanDocumentContext $ctx): void {}
    public function PreDone(InkspanDocumentContext $ctx): void {}
    public function PreUnexpected(InkspanDocumentContext $ctx): void {}
}
