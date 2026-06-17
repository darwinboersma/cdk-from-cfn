package com.myorg;

import software.constructs.Construct;

import java.util.*;
import software.amazon.awscdk.CfnMapping;
import software.amazon.awscdk.CfnTag;
import software.amazon.awscdk.Stack;
import software.amazon.awscdk.StackProps;

import software.amazon.awscdk.*;
import software.amazon.awscdk.services.sns.*;

class ConditionSubConstruct extends Construct {
    public ConditionSubConstruct(final Construct scope, final String id) {
        this(scope, id, null);
    }

    public ConditionSubConstruct(final Construct scope, final String id,
            String envName) {
        super(scope, id);


        Boolean hasCustomSuffix = (envName + "-suffix").equals("prod-suffix");

        Optional<CfnTopic> topic = hasCustomSuffix ? Optional.of(CfnTopic.Builder.create(this, "Topic")
                .build()) : Optional.empty();

    }
}
