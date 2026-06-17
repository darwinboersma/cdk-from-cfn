package com.myorg;

import software.constructs.Construct;

import java.util.*;
import software.amazon.awscdk.CfnMapping;
import software.amazon.awscdk.CfnTag;
import software.amazon.awscdk.Stack;
import software.amazon.awscdk.StackProps;

import software.amazon.awscdk.*;
import software.amazon.awscdk.services.sns.*;

class ConditionSubStack extends Stack {
    public ConditionSubStack(final Construct scope, final String id) {
        super(scope, id, null);
    }

    public ConditionSubStack(final Construct scope, final String id, final StackProps props) {
        this(scope, id, props, null);
    }

    public ConditionSubStack(final Construct scope, final String id, final StackProps props,
            String envName) {
        super(scope, id, props);


        Boolean hasCustomSuffix = (envName + "-suffix").equals("prod-suffix");

        Optional<CfnTopic> topic = hasCustomSuffix ? Optional.of(CfnTopic.Builder.create(this, "Topic")
                .build()) : Optional.empty();

    }
}
