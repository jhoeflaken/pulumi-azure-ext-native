// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azureext.certificate;

import com.pulumi.azureext.Utilities;
import com.pulumi.azureext.certificate.KeyVaultCertificateArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import javax.annotation.Nullable;

/**
 * An Azure key vault certificate.
 * 
 */
@ResourceType(type="azure-ext:certificate:KeyVaultCertificate")
public class KeyVaultCertificate extends com.pulumi.resources.CustomResource {
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public KeyVaultCertificate(String name) {
        this(name, KeyVaultCertificateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KeyVaultCertificate(String name, @Nullable KeyVaultCertificateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KeyVaultCertificate(String name, @Nullable KeyVaultCertificateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azure-ext:certificate:KeyVaultCertificate", name, args == null ? KeyVaultCertificateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private KeyVaultCertificate(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azure-ext:certificate:KeyVaultCertificate", name, null, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static KeyVaultCertificate get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new KeyVaultCertificate(name, id, options);
    }
}