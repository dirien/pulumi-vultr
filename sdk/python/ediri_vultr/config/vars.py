# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('vultr')


class _ExportableConfig(types.ModuleType):
    @property
    def api_key(self) -> Optional[str]:
        """
        The API Key that allows interaction with the API
        """
        return __config__.get('apiKey') or _utilities.get_env('VULTR_API_KEY')

    @property
    def rate_limit(self) -> int:
        """
        Allows users to set the speed of API calls to work with the Vultr Rate Limit
        """
        return __config__.get_int('rateLimit') or 500

    @property
    def retry_limit(self) -> int:
        """
        Allows users to set the maximum number of retries allowed for a failed API call.
        """
        return __config__.get_int('retryLimit') or 3

