<?php
declare(strict_types=1);

// InkspanDocument SDK

require_once __DIR__ . '/utility/struct/Struct.php';
require_once __DIR__ . '/core/UtilityType.php';
require_once __DIR__ . '/core/Spec.php';
require_once __DIR__ . '/core/Helpers.php';

// Load utility registration
require_once __DIR__ . '/utility/Register.php';

// Load config and features
require_once __DIR__ . '/config.php';
require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/features.php';

use Voxgig\Struct\Struct;

// Features record diagnostic state on the client as dynamic properties
// (_retry, _cache, _metrics, ...); allow them explicitly (PHP 8.2+
// deprecates implicit dynamic properties).
#[\AllowDynamicProperties]
class InkspanDocumentSDK implements \JsonSerializable
{
    public string $mode;
    public array $features;
    public ?array $options;

    private $_utility;
    private $_rootctx;

    public function __construct(array $options = [])
    {
        $this->mode = "live";
        $this->features = [];
        $this->options = null;

        $utility = new InkspanDocumentUtility();
        $this->_utility = $utility;

        $config = InkspanDocumentConfig::shared_config();

        $this->_rootctx = ($utility->make_context)([
            "client" => $this,
            "utility" => $utility,
            "config" => $config,
            "options" => $options ?? [],
            "shared" => [],
        ], null);

        $this->options = ($utility->make_options)($this->_rootctx);

        if (Struct::getpath($this->options, "feature.test.active") === true) {
            $this->mode = "test";
        }

        $this->_rootctx->options = $this->options;

        // Feature INSTANCES supplied at construction (the station adopt
        // path) are read from the RAW construction options - extend is
        // consumed exactly once, here; make_options strips it from the
        // processed map so options_map() stays clean data.
        $extend_val = is_array($options["extend"] ?? null) ? $options["extend"] : [];

        // Add features in the resolved order (make_options puts an explicit
        // list order first, else defaults to test-first). Ordering matters: the
        // `test` feature installs the base mock transport and the transport
        // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is
        // current, so `test` must be added before them to sit at the base.
        $feature_opts = InkspanDocumentHelpers::to_map(Struct::getprop($this->options, "feature"));
        if ($feature_opts) {
            $featureorder = Struct::getpath($this->options, "__derived__.featureorder");
            if (is_array($featureorder)) {
                foreach ($featureorder as $fname) {
                    $fopts = InkspanDocumentHelpers::to_map($feature_opts[$fname] ?? null);
                    if ($fopts && isset($fopts["active"]) && $fopts["active"] === true) {
                        // An active name with no generated feature class is
                        // legal when an extend-supplied instance carries that
                        // name (station's adopt path): the instance is added
                        // below, positioned by its own __after__ entry, so
                        // skip it here rather than add a BaseFeature stray
                        // that would silently shift feature positions.
                        if (!InkspanDocumentFeatures::has_feature($fname)) {
                            foreach ($extend_val as $ef) {
                                if (is_object($ef) && method_exists($ef, 'get_name')
                                    && $fname === $ef->get_name()) {
                                    continue 2;
                                }
                            }
                        }
                        ($utility->feature_add)($this->_rootctx, InkspanDocumentFeatures::make_feature($fname));
                    }
                }
            }
        }

        // Add extension features.
        foreach ($extend_val as $f) {
            if (is_object($f) && method_exists($f, 'get_name')) {
                ($utility->feature_add)($this->_rootctx, $f);
            }
        }

        // Initialize features.
        foreach ($this->features as $f) {
            ($utility->feature_init)($this->_rootctx, $f);
        }

        ($utility->feature_hook)($this->_rootctx, "PostConstruct");
    }

    public function options_map(): array
    {
        $out = Struct::clone($this->options);
        return is_array($out) ? $out : [];
    }

    // The options hold the credential, so the default print and json form
    // name the client and nothing more; options_map() is the way to read
    // them back.
    public function jsonSerialize(): array
    {
        return ['name' => 'InkspanDocument'];
    }

    public function __debugInfo(): array
    {
        return $this->jsonSerialize();
    }

    public function get_utility()
    {
        return InkspanDocumentUtility::copy($this->_utility);
    }

    public function get_root_ctx()
    {
        return $this->_rootctx;
    }

    public function prepare(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;
        $fetchargs = $fetchargs ?? [];

        $ctrl = InkspanDocumentHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "prepare",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $opts = $this->options;
        $path = Struct::getprop($fetchargs, "path") ?? "";
        $path = is_string($path) ? $path : "";
        $method_val = Struct::getprop($fetchargs, "method") ?? "GET";
        $method_val = strtoupper(is_string($method_val) && '' !== $method_val ? $method_val : "GET");
        $allow_method = Struct::getpath($opts, "allow.method");
        if (!InkspanDocumentPrepareMethod::allowed($allow_method, $method_val)) {
            return ($utility->make_error)($ctx, $ctx->make_error("spec_method_allow",
                "Method \"" . $method_val . "\" not allowed by SDK option allow.method value: \"" .
                (is_string($allow_method) ? $allow_method : "") . "\""));
        }
        $params = InkspanDocumentHelpers::to_map(Struct::getprop($fetchargs, "params")) ?? [];
        $query = InkspanDocumentHelpers::to_map(Struct::getprop($fetchargs, "query")) ?? [];
        $headers = ($utility->prepare_headers)($ctx);

        $base = Struct::getprop($opts, "base") ?? "";
        $base = is_string($base) ? $base : "";
        $prefix = Struct::getprop($opts, "prefix") ?? "";
        $prefix = is_string($prefix) ? $prefix : "";
        $suffix = Struct::getprop($opts, "suffix") ?? "";
        $suffix = is_string($suffix) ? $suffix : "";

        $ctx->spec = new InkspanDocumentSpec([
            "base" => $base, "prefix" => $prefix, "suffix" => $suffix,
            "path" => $path, "method" => $method_val,
            "params" => $params, "query" => $query, "headers" => $headers,
            "body" => Struct::getprop($fetchargs, "body"),
            "step" => "start",
        ]);

        // Merge user-provided headers.
        $uh = Struct::getprop($fetchargs, "headers");
        if (is_array($uh)) {
            foreach ($uh as $k => $v) {
                $ctx->spec->headers[$k] = $v;
            }
        }

        [$_, $err] = ($utility->prepare_auth)($ctx);
        if ($err) {
            return ($utility->make_error)($ctx, $err);
        }

        [$fetchdef, $fd_err] = ($utility->make_fetch_def)($ctx);
        if ($fd_err) {
            return ($utility->make_error)($ctx, $fd_err);
        }
        return $fetchdef;
    }

    // Raw endpoint access is operator-controllable, like every entity op.
    // Blocking it means denying BOTH the 'direct' and 'graphql' tokens,
    // since either one reaches the same endpoint.
    public function direct(array $fetchargs = []): mixed
    {
        if (!$this->op_allowed("direct")) {
            return $this->op_denied("direct");
        }

        return $this->raw_request($fetchargs);
    }

    // Is this raw-access op permitted by the SDK's allow.op option?
    private function op_allowed(string $op): bool
    {
        return InkspanDocumentPrepareMethod::allowed(Struct::getpath($this->options, "allow.op"), $op);
    }

    private function op_denied(string $op): array
    {
        $allow_op = Struct::getpath($this->options, "allow.op");
        return [
            "ok" => false,
            "err" => new InkspanDocumentError($op . "_allow",
                "InkspanDocumentSDK: " . $op . ": operation not allowed by" .
                " SDK option allow.op value: \"" . (string)$allow_op . "\""),
        ];
    }

    // Ungated request path shared by direct and graphql, each of which
    // checks its own allow.op token first. Private, rather than a flag on
    // fetchargs: a caller-supplied marker would let anyone opt straight back
    // out of the gate by passing it.
    private function raw_request(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;

        // direct() is the raw-HTTP escape hatch: it never throws, it returns
        // an {ok, err, ...} dict. prepare() now raises on error, so catch it
        // and surface the failure through the dict instead.
        try {
            $fetchdef = $this->prepare($fetchargs);
        } catch (\Throwable $err) {
            return ["ok" => false, "err" => $err];
        }

        $fetchargs = $fetchargs ?? [];
        $ctrl = InkspanDocumentHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "direct",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $url = $fetchdef["url"] ?? "";
        [$fetched, $fetch_err] = ($utility->fetcher)($ctx, $url, $fetchdef);

        if ($fetch_err) {
            return ["ok" => false, "err" => ($utility->clean)($ctx, $fetch_err)];
        }

        if ($fetched === null) {
            return [
                "ok" => false,
                "err" => $ctx->make_error("direct_no_response", "response: undefined"),
            ];
        }

        if (is_array($fetched)) {
            $status = InkspanDocumentHelpers::to_int(Struct::getprop($fetched, "status"));
            $headers = Struct::getprop($fetched, "headers") ?? [];

            // No-body responses (204, 304) and explicit zero content-length
            // must skip JSON parsing — calling json() on an empty body errors.
            $content_length = is_array($headers) ? ($headers["content-length"] ?? null) : null;
            $no_body = $status === 204 || $status === 304 || (string)$content_length === "0";

            $json_data = null;
            $body_err = null;
            if (!$no_body) {
                $jf = Struct::getprop($fetched, "json");
                if (is_callable($jf)) {
                    try {
                        $json_data = $jf();
                    } catch (\Throwable $e) {
                        // Non-JSON body — leave data null but keep status/ok.
                        $json_data = null;
                    }
                }
                if (true === Struct::getprop($fetched, "unreadable")) {
                    $failed = ($status >= 200 && $status < 300) ? null : $ctx->make_error(
                        "request_status",
                        "request: {$status}: " . (string)Struct::getprop($fetched, "statusText"));
                    $sent = $fetchdef["headers"] ?? [];
                    if (InkspanDocumentFetcher::usesDefault($ctx)) {
                        $sent = InkspanDocumentFetcher::sentHeaders($sent);
                    }
                    $body_err = InkspanDocumentResultBody::unreadable($ctx, $status, $headers,
                        Struct::getprop($fetched, "body"), $sent, $failed);
                }
            }

            $out = [
                "ok" => null === $body_err && $status >= 200 && $status < 300,
                "status" => $status,
                "headers" => Struct::getprop($fetched, "headers"),
                "data" => $json_data,
            ];
            if (null !== $body_err) {
                $out["err"] = ($utility->clean)($ctx, $body_err);
            }
            return $out;
        }

        return [
            "ok" => false,
            "err" => $ctx->make_error("direct_invalid", "invalid response type"),
        ];
    }

    // Raw GraphQL access: the pressure valve that makes the generated
    // surface's deliberate omissions (per-call selection sets, typed filter
    // builders, batching, subscriptions) livable — the whole schema stays
    // reachable.
    //
    // Thin wrapper over the same prepare/fetch path direct uses, with the
    // one thing raw direct cannot do for GraphQL: a GraphQL failure rides
    // HTTP 200 as a top-level `errors` array, so status alone would report
    // a failed query as ok.
    //
    // NOTE: like direct, this bypasses the feature pipeline — no retry,
    // ratelimit or paging features apply.
    public function graphql(string $query, ?array $variables = null, ?array $ctrl = null): mixed
    {
        if (!$this->op_allowed("graphql")) {
            return $this->op_denied("graphql");
        }

        $res = $this->raw_request([
            "method" => "POST",
            "headers" => ["content-type" => "application/json"],
            "body" => ["query" => $query, "variables" => $variables ?? []],
            "ctrl" => $ctrl ?? [],
        ]);

        if (!is_array($res)) {
            return $res;
        }

        // Errors are read BEFORE any status check: a GraphQL parse or
        // validation failure comes back as HTTP 400 carrying the standard
        // { errors: [...] } body, and the raw path represents a non-2xx as
        // ok:false with no err — so returning early on status would discard
        // the server's own diagnostics, which are the only useful part of
        // that response.
        $errors = Struct::getpath($res, "data.errors");

        if (is_array($errors) && 0 < count($errors)) {
            $first = is_array($errors[0]) ? $errors[0] : [];
            $msg = $first["message"] ?? "";
            if (!is_string($msg) || "" === $msg) {
                $msg = "graphql error";
            }
            $res["ok"] = false;
            $res["err"] = new InkspanDocumentError("graphql_error",
                "InkspanDocumentSDK: graphql: " . $msg);
            $res["graphql"] = $errors;
        }

        return $res;
    }


    private $_checkout = null;

    // Canonical facade: $client->Checkout()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->checkout()
    // resolves here too.
    public function Checkout($data = null)
    {
        require_once __DIR__ . '/entity/checkout_entity.php';
        if ($data === null) {
            if ($this->_checkout === null) {
                $this->_checkout = new CheckoutEntity($this, null);
            }
            return $this->_checkout;
        }
        return new CheckoutEntity($this, $data);
    }


    private $_compress = null;

    // Canonical facade: $client->Compress()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->compress()
    // resolves here too.
    public function Compress($data = null)
    {
        require_once __DIR__ . '/entity/compress_entity.php';
        if ($data === null) {
            if ($this->_compress === null) {
                $this->_compress = new CompressEntity($this, null);
            }
            return $this->_compress;
        }
        return new CompressEntity($this, $data);
    }


    private $_fill_form = null;

    // Canonical facade: $client->FillForm()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->fill_form()
    // resolves here too.
    public function FillForm($data = null)
    {
        require_once __DIR__ . '/entity/fill_form_entity.php';
        if ($data === null) {
            if ($this->_fill_form === null) {
                $this->_fill_form = new FillFormEntity($this, null);
            }
            return $this->_fill_form;
        }
        return new FillFormEntity($this, $data);
    }


    private $_health = null;

    // Canonical facade: $client->Health()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->health()
    // resolves here too.
    public function Health($data = null)
    {
        require_once __DIR__ . '/entity/health_entity.php';
        if ($data === null) {
            if ($this->_health === null) {
                $this->_health = new HealthEntity($this, null);
            }
            return $this->_health;
        }
        return new HealthEntity($this, $data);
    }


    private $_html = null;

    // Canonical facade: $client->Html()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->html()
    // resolves here too.
    public function Html($data = null)
    {
        require_once __DIR__ . '/entity/html_entity.php';
        if ($data === null) {
            if ($this->_html === null) {
                $this->_html = new HtmlEntity($this, null);
            }
            return $this->_html;
        }
        return new HtmlEntity($this, $data);
    }


    private $_inbound = null;

    // Canonical facade: $client->Inbound()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->inbound()
    // resolves here too.
    public function Inbound($data = null)
    {
        require_once __DIR__ . '/entity/inbound_entity.php';
        if ($data === null) {
            if ($this->_inbound === null) {
                $this->_inbound = new InboundEntity($this, null);
            }
            return $this->_inbound;
        }
        return new InboundEntity($this, $data);
    }


    private $_markdown = null;

    // Canonical facade: $client->Markdown()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->markdown()
    // resolves here too.
    public function Markdown($data = null)
    {
        require_once __DIR__ . '/entity/markdown_entity.php';
        if ($data === null) {
            if ($this->_markdown === null) {
                $this->_markdown = new MarkdownEntity($this, null);
            }
            return $this->_markdown;
        }
        return new MarkdownEntity($this, $data);
    }


    private $_men = null;

    // Canonical facade: $client->Men()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->men()
    // resolves here too.
    public function Men($data = null)
    {
        require_once __DIR__ . '/entity/men_entity.php';
        if ($data === null) {
            if ($this->_men === null) {
                $this->_men = new MenEntity($this, null);
            }
            return $this->_men;
        }
        return new MenEntity($this, $data);
    }


    private $_merge = null;

    // Canonical facade: $client->Merge()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->merge()
    // resolves here too.
    public function Merge($data = null)
    {
        require_once __DIR__ . '/entity/merge_entity.php';
        if ($data === null) {
            if ($this->_merge === null) {
                $this->_merge = new MergeEntity($this, null);
            }
            return $this->_merge;
        }
        return new MergeEntity($this, $data);
    }


    private $_portal = null;

    // Canonical facade: $client->Portal()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->portal()
    // resolves here too.
    public function Portal($data = null)
    {
        require_once __DIR__ . '/entity/portal_entity.php';
        if ($data === null) {
            if ($this->_portal === null) {
                $this->_portal = new PortalEntity($this, null);
            }
            return $this->_portal;
        }
        return new PortalEntity($this, $data);
    }


    private $_root_get = null;

    // Canonical facade: $client->RootGet()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->root_get()
    // resolves here too.
    public function RootGet($data = null)
    {
        require_once __DIR__ . '/entity/root_get_entity.php';
        if ($data === null) {
            if ($this->_root_get === null) {
                $this->_root_get = new RootGetEntity($this, null);
            }
            return $this->_root_get;
        }
        return new RootGetEntity($this, $data);
    }


    private $_signup = null;

    // Canonical facade: $client->Signup()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->signup()
    // resolves here too.
    public function Signup($data = null)
    {
        require_once __DIR__ . '/entity/signup_entity.php';
        if ($data === null) {
            if ($this->_signup === null) {
                $this->_signup = new SignupEntity($this, null);
            }
            return $this->_signup;
        }
        return new SignupEntity($this, $data);
    }


    private $_split = null;

    // Canonical facade: $client->Split()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->split()
    // resolves here too.
    public function Split($data = null)
    {
        require_once __DIR__ . '/entity/split_entity.php';
        if ($data === null) {
            if ($this->_split === null) {
                $this->_split = new SplitEntity($this, null);
            }
            return $this->_split;
        }
        return new SplitEntity($this, $data);
    }


    private $_table = null;

    // Canonical facade: $client->Table()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->table()
    // resolves here too.
    public function Table($data = null)
    {
        require_once __DIR__ . '/entity/table_entity.php';
        if ($data === null) {
            if ($this->_table === null) {
                $this->_table = new TableEntity($this, null);
            }
            return $this->_table;
        }
        return new TableEntity($this, $data);
    }


    private $_template = null;

    // Canonical facade: $client->Template()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->template()
    // resolves here too.
    public function Template($data = null)
    {
        require_once __DIR__ . '/entity/template_entity.php';
        if ($data === null) {
            if ($this->_template === null) {
                $this->_template = new TemplateEntity($this, null);
            }
            return $this->_template;
        }
        return new TemplateEntity($this, $data);
    }


    private $_text = null;

    // Canonical facade: $client->Text()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->text()
    // resolves here too.
    public function Text($data = null)
    {
        require_once __DIR__ . '/entity/text_entity.php';
        if ($data === null) {
            if ($this->_text === null) {
                $this->_text = new TextEntity($this, null);
            }
            return $this->_text;
        }
        return new TextEntity($this, $data);
    }


    private $_to_image = null;

    // Canonical facade: $client->ToImage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->to_image()
    // resolves here too.
    public function ToImage($data = null)
    {
        require_once __DIR__ . '/entity/to_image_entity.php';
        if ($data === null) {
            if ($this->_to_image === null) {
                $this->_to_image = new ToImageEntity($this, null);
            }
            return $this->_to_image;
        }
        return new ToImageEntity($this, $data);
    }


    private $_translate = null;

    // Canonical facade: $client->Translate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->translate()
    // resolves here too.
    public function Translate($data = null)
    {
        require_once __DIR__ . '/entity/translate_entity.php';
        if ($data === null) {
            if ($this->_translate === null) {
                $this->_translate = new TranslateEntity($this, null);
            }
            return $this->_translate;
        }
        return new TranslateEntity($this, $data);
    }


    private $_watermark = null;

    // Canonical facade: $client->Watermark()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->watermark()
    // resolves here too.
    public function Watermark($data = null)
    {
        require_once __DIR__ . '/entity/watermark_entity.php';
        if ($data === null) {
            if ($this->_watermark === null) {
                $this->_watermark = new WatermarkEntity($this, null);
            }
            return $this->_watermark;
        }
        return new WatermarkEntity($this, $data);
    }


    private $_webhook = null;

    // Canonical facade: $client->Webhook()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->webhook()
    // resolves here too.
    public function Webhook($data = null)
    {
        require_once __DIR__ . '/entity/webhook_entity.php';
        if ($data === null) {
            if ($this->_webhook === null) {
                $this->_webhook = new WebhookEntity($this, null);
            }
            return $this->_webhook;
        }
        return new WebhookEntity($this, $data);
    }



    public static function test(?array $testopts = null, ?array $sdkopts = null): self
    {
        $sdkopts = $sdkopts ?? [];
        $sdkopts = Struct::clone($sdkopts);
        $sdkopts = is_array($sdkopts) ? $sdkopts : [];

        $testopts = $testopts ?? [];
        $testopts = Struct::clone($testopts);
        $testopts = is_array($testopts) ? $testopts : [];
        $testopts["active"] = true;

        if (!isset($sdkopts["feature"])) {
            $sdkopts["feature"] = [];
        }
        $sdkopts["feature"]["test"] = $testopts;

        $sdk = new InkspanDocumentSDK($sdkopts);
        $sdk->mode = "test";
        return $sdk;
    }
}
