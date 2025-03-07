import { intOptions } from './../node_modules/aws-cdk-lib/node_modules/yaml/types.d';
import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import { RestApi, LambdaIntegration } from "aws-cdk-lib/aws-apigateway";

export class AwsLambdaGoStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const lambdaFunction = new lambda.Function(this, 'TutorialLambda', {
      code: lambda.Code.fromAsset('lambda'),
      handler: 'main',
      runtime: lambda.Runtime.PROVIDED_AL2023,
    });
    
    const gateway = new RestApi(this, 'TutorialGateway', {
      defaultCorsPreflightOptions: {
        allowOrigins: ['*'],
        allowMethods: ['GET', 'POST', 'PUT', 'DELETE', 'OPTIONS'],
      }
    });

    const integration = new LambdaIntegration(lambdaFunction);
    const testResource = gateway.root.addResource('test');
    testResource.addMethod('GET', integration, {});
  }
}
