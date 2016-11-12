<?php

namespace Phuse\Framework\Module\{{.Name}};

use Phuse\Framework\Module\Module\AbstractModule;
use Phuse\Framework\Module\Processor\ProcessorInterface;

class {{.Name}}Module extends AbstractModule implements {{.Name}}ModuleInterface
{
    /** @var {{.Name}}DataContainer */
    private $dataContainer;

    /** @var {{.Name}}HandlerContainer */
    private $handlerContainer;

    /** @var {{.Name}}ProcessorContainer */
    private $processorContainer;

    /** @var string */
    private $moduleName = '{{.Name}}';


    /**
     * @return bool
     */
    public function init()
    {
        return true;
    }

    /**
     * @return string
     */
    public function getName()
    {
        return lcfirst($this->moduleName);
    }

    /**
     * @return {{.Name}}DataContainer
     */
    public function getDataContainer()
    {
        if (!$this->dataContainer)
        {
            $this->dataContainer = (new {{.Name}}DataContainer())
                ->setProduct($this->getProduct());
        }

        return $this->dataContainer;
    }

    /**
     * @return {{.Name}}HandlerContainer
     */
    public function getHandlerContainer()
    {
        if (!$this->handlerContainer)
        {
            $this->handlerContainer = (new {{.Name}}HandlerContainer())
                ->setProduct($this->getProduct())
                ->setProvider($this->getProvider());
        }

        return $this->handlerContainer;
    }

    /**
     * @return {{.Name}}ProcessorContainer
     */
    public function getProcessorContainer()
    {
        if (!$this->processorContainer)
        {
            $this->processorContainer = (new {{.Name}}ProcessorContainer())
                ->setProduct($this->getProduct())
                ->setProvider($this->getProvider());
        }

        return $this->processorContainer;
    }

    /**
     * @return ProcessorInterface[]
     */
    public function getTasks()
    {
        return $this
            ->getProcessorContainer()
            ->getProcessorSetupHandler()
            ->getProcessors();
    }
}
